package research

import "obst"

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

func EvaluateSeq(cfg *Config, tree *obst.Tree) {
	N := cfg.N

	for i := N; i >= 0; i-- {
		for j := i; j <= N; j++ {
			root, cost := tree.Calculate(i, j)
			tree.Root[i][j] = root
			tree.Cost[i][j] = cost
		}
	}
	return
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Dependency ...
//  	Data set for chunk calculation & synchronization
type Dependency struct {
	PreSet  [2]chan int
	PostSet [2]chan int
}

// Wait ...
func (rcv *Dependency) Wait() {
	if rcv.PreSet[0] != nil {
		<-rcv.PreSet[0]
	}
	if rcv.PreSet[1] != nil {
		<-rcv.PreSet[1]
	}
}

// Notify ...
func (rcv *Dependency) Notify() {
	if rcv.PostSet[0] != nil {
		rcv.PostSet[0] <- 1
	}
	if rcv.PostSet[1] != nil {
		rcv.PostSet[1] <- 1
	}
}

func NewDepMatrix(row int, col int) (mat [][]Dependency) {
	size := row * col
	arr := make([]Dependency, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}

func Chunk(tree *obst.Tree, i int, j int, width int, dep *Dependency) {
	dep.Wait()

	for row := i + width - 1; row >= i; row-- {
		for col := j; col < j+width; col++ {
			root, cost := tree.Calculate(row, col)
			tree.Root[row][col] = root
			tree.Cost[row][col] = cost
		}
	}

	dep.Notify()
}

// EvaluatePar ...
//  	Parallel Chunk processing
func EvaluatePar(cfg *Config, tree *obst.Tree) {
	N, VP := cfg.N, cfg.VP
	Hor, Ver := 0, 1 // index alias

	deps := NewDepMatrix(VP, VP)

	// Shared Channels for sync
	chs := new(Channels)
	chs.Init(VP - 1) // Allocate channels

	// Last chunk's dependency
	{
		x, y := 0, VP-1
		deps[x][y].PostSet[Hor] = chs.Finish
	}

	// Horizontal dependecy
	for x := 0; x < VP-1; x++ {
		for y := 0; y < VP-1; y++ {
			relay := chs.H[x][y]
			// Chunk[x][y] -> H[x][y] -> Chunk[x][y+1]
			deps[x][y].PostSet[Hor] = relay  // Chunk[x][y] -> H[x][y]
			deps[x][y+1].PreSet[Hor] = relay // H[x][y] -> Chunk[x][y+1]
		}
	}

	// Vertical dependecy
	for x := 1; x < VP; x++ {
		for y := 1; y < VP; y++ {
			relay := chs.V[x-1][y-1]
			// Chunk[x][y] -> V[x-1][y-1] -> Chunk[x-1][y]
			deps[x][y].PostSet[Ver] = relay  // Chunk[x][y] -> V[x-1][y-1]
			deps[x-1][y].PreSet[Ver] = relay //V[x-1][y-1] -> Chunk[x-1][y]
		}
	}

	// width of each chunk
	width := N / VP

	// Delegate chunks to goroutine
	for x, i := 0, 0; x < VP; x++ {
		for y, j := x, 1; y < VP; y++ {

			dep := deps[x][y] // dependency of the chunk
			go Chunk(tree, i, j, width, &dep)

			j += width
		}
		i += width
	}
	<-chs.Finish
	return
}
