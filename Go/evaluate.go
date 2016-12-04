package main

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
func EvaluatePar(cfg *Config, tree *obst.Tree, chs *Channels) {
	N, VP := cfg.N, cfg.VP
	H, V := 0, 1

	deps := NewDepMatrix(VP, VP)

	// Last chunk's dependency
	{
		x, y := 0, VP-1
		// fmt.Printf("Last chunk : %2d, %2d \n", x, y)
		deps[x][y].PostSet[H] = chs.Finish
	}

	// Horizontal dependecy
	for x := 0; x < VP-1; x++ {
		for y := 0; y < VP-1; y++ {
			// fmt.Printf("H-Relation : %2d, %2d \n", x, y)

			relay := chs.H[x][y]
			// Chunk[x][y] -> H[x][y] -> Chunk[x][y+1]
			deps[x][y].PostSet[H] = relay  // Chunk[x][y] -> H[x][y]
			deps[x][y+1].PreSet[H] = relay // H[x][y] -> Chunk[x][y+1]
		}
	}

	// Vertical dependecy
	for x := 1; x < VP; x++ {
		for y := 1; y < VP; y++ {
			// fmt.Printf("V-Relation : %2d, %2d \n", x, y)

			relay := chs.V[x-1][y-1]
			// Chunk[x][y] -> V[x-1][y-1] -> Chunk[x-1][y]
			deps[x][y].PostSet[V] = relay  // Chunk[x][y] -> V[x-1][y-1]
			deps[x-1][y].PreSet[V] = relay //V[x-1][y-1] -> Chunk[x-1][y]
		}
	}

	// Delegate chunks to goroutine
	for x, i := 0, 0; x < VP; x++ {
		for y, j := x, 1; y < VP; y++ {
			dep := deps[x][y] // dependency of the chunk
			width := N / VP   // chunk's width

			// fmt.Printf("Chunk : x:%2d, y:%2d, i:%2d,j:%2d, w:%3d \n", x, y, i, j, width)
			go Chunk(tree, i, j, width, &dep)

			j += VP
		}
		i += VP
	}
	<-chs.Finish
	return
}
