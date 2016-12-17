// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
// 	File	: evaluate.go
//  Author	: Park Dong Ha ( luncliff@gmail.com )
//  Updated	: 2016/12/17
//
//  Note	:
//      Evaluate Optimal Binary Search Tree problem.
//      - `EvaluateSeq` : Sequential processing
//      - `EvaluatePar` : Parallel processing with Goroutine
//
//  Note     :
//      The sub-problems are grouped to a chunk, and each goroutine
//		will process given single chunk.
//
//  Concept  :
//      - Main Problem :
//          Evaluating a `Tree`.
//          For parallel processing, it is divided into sub-problems.
//      - Sub-problem :
//          Calculating `root` and `cost` with given vertices.
//      - Chunk
//          To process efficiently, sub-problems are *chunked*.
//          The size of chunk can be small so each chunks
//			can be mapped to sub-problem in 1:1 relation. (VP==N)
//          Or, it can be big to reduce synchronization overhead.
//          (VP << N)
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package research

import "obst"

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// EvaluateSeq ...
//  	Sequential Evaluation
func EvaluateSeq(cfg *Config, tree *obst.Tree) {
	N := cfg.N

	// loop : sequential processing
	//		[ . . . . ]
	//		[ 9 . . . ]
	//		[ 5 6 . . ]
	//	->	[ 1 2 3 4 ]
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
//  	Set of channels for chunk synchronization
type Dependency struct {
	PreSet  [2]chan int
	PostSet [2]chan int
}

// Wait ...
//		Wait for pre-set's notification
func (rcv *Dependency) Wait() {
	if rcv.PreSet[0] != nil {
		<-rcv.PreSet[0]
	}
	if rcv.PreSet[1] != nil {
		<-rcv.PreSet[1]
	}
}

// Notify ...
//  	Notify to post-set using channels
func (rcv *Dependency) Notify() {
	if rcv.PostSet[0] != nil {
		rcv.PostSet[0] <- 1
	}
	if rcv.PostSet[1] != nil {
		rcv.PostSet[1] <- 1
	}
}

// NewDepMatrix ...
//  	Create a matrix of Dependency object
func NewDepMatrix(row int, col int) (mat [][]Dependency) {
	size := row * col
	arr := make([]Dependency, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}

// Chunk ...
func Chunk(tree *obst.Tree, i int, j int, width int, dep *Dependency) {
	dep.Wait() // Wait for preset...

	// loop : sequential processing
	//		[ . . . . ]
	//		[ 9 . . . ]
	//		[ 5 6 . . ]
	//	->	[ 1 2 3 4 ]
	for row := i + width - 1; row >= i; row-- {
		for col := j; col < j+width; col++ {
			root, cost := tree.Calculate(row, col)
			tree.Root[row][col] = root
			tree.Cost[row][col] = cost
		}
	}

	dep.Notify() // Notify to postset...
}

// EvaluatePar ...
//  	Parallel processing
func EvaluatePar(cfg *Config, tree *obst.Tree) {
	N, VP := cfg.N, cfg.VP

	// Dependency Matrix for synchronization
	deps := NewDepMatrix(VP, VP)

	// Shared Channels for sync
	chs := new(Channels)
	chs.Init(VP - 1) // Allocate channels

	H, V := 0, 1 // index alias

	// Last chunk's dependency
	// This chunk notifies the end of processing
	{
		x, y := 0, VP-1
		deps[x][y].PostSet[H] = chs.Finish
	}

	// Horizontal dependecy
	for x := 0; x < VP-1; x++ {
		for y := 0; y < VP-1; y++ {
			relay := chs.H[x][y]
			// Chunk[x][y] -> H[x][y] -> Chunk[x][y+1]
			deps[x][y].PostSet[H] = relay  // Chunk[x][y] -> H[x][y]
			deps[x][y+1].PreSet[H] = relay // H[x][y] -> Chunk[x][y+1]
		}
	}

	// Vertical dependecy
	for x := 1; x < VP; x++ {
		for y := 1; y < VP; y++ {
			relay := chs.V[x-1][y-1]
			// Chunk[x][y] -> V[x-1][y-1] -> Chunk[x-1][y]
			deps[x][y].PostSet[V] = relay  // Chunk[x][y] -> V[x-1][y-1]
			deps[x-1][y].PreSet[V] = relay // V[x-1][y-1] -> Chunk[x-1][y]
		}
	}

	// width of each chunk
	width := N / VP

	// Delegate chunks to goroutine
	for x, i := 0, 0; x < VP; x++ {
		for y, j := x, 1; y < VP; y++ {

			// Make goroutine to process chunk
			// aware with dependency
			go Chunk(tree, i, j, width, &deps[x][y])

			j += width // jump column
		}
		i += width // jump row
	}

	// When last goroutine done. It will notify finish.
	<-chs.Finish
	return
}
