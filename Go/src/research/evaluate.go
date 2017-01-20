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

import (
	"obst"
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// EvaluateSeq ...
//  	Sequential Evaluation
func EvaluateSeq(tree *obst.Tree) {
	N := tree.Size()

	// loop : sequential processing
	//		[ 7 8 . . ]
	//		[ . 4 5 6 ]
	//		[ . . 2 3 ]
	//	->	[ . . . 1 ]
	for i := N; i >= 0; i-- {
		for j := i; j <= N; j++ {
			root, cost := tree.Calculate(i, j)
			*tree.Root.At(i, j) = root
			*tree.Cost.At(i, j) = cost
		}
	}
	return
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// index alias
const (
	H int = 0 // Horizontal
	V int = 1 // Vertical
)

// Dependency ...
//  	Set of channels for chunk synchronization
type Dependency struct {
	PreSet  [2]chan int
	PostSet [2]chan int
}

// Wait ...
//		Wait for pre-set's notification
// 		The order is related to EvaluatePar's spawning loop
func (rcv *Dependency) Wait() {
	// Wait vertical channel first
	if rcv.PreSet[V] != nil {
		<-rcv.PreSet[V]
	}
	// Wait horizontal channel
	if rcv.PreSet[H] != nil {
		<-rcv.PreSet[H]
	}
}

// Notify ...
//  	Notify to post-set using channels
func (rcv *Dependency) Notify() {
	// Notify vertical channel first
	if rcv.PostSet[V] != nil {
		rcv.PostSet[V] <- 1
	}
	// Notify horizontal channel
	if rcv.PostSet[H] != nil {
		rcv.PostSet[H] <- 1
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
func Chunk(
	tree *obst.Tree,
	i int, j int, width int, dep *Dependency) {

	// 1. Wait for pre-set...
	dep.Wait()

	// 2. Sequential processing
	//		[ . . . . ]
	//		[ 9 . . . ]
	//		[ 5 6 7 8 ]
	//	->	[ 1 2 3 4 ]
	for row := i - 1 + width; i <= row; row-- {
		for col := j; col < j+width; col++ {

			root, cost := tree.Calculate(row, col)
			*tree.Root.At(row, col) = root
			*tree.Cost.At(row, col) = cost
		}
	}

	// 3. Notify to post-set...
	dep.Notify()
}

// EvaluatePar ...
//  	Parallel processing
func EvaluatePar(tree *obst.Tree, VP int) {
	N := tree.Size()

	// Dependency Matrix for synchronization
	deps := NewDepMatrix(VP, VP)
	// Shared Channels for sync
	chs := new(Channels)
	chs.Init(VP - 1) // Allocate channels

	// Final chunk's dependency
	// This chunk notifies the end of processing
	//    [ - - - F ]
	//    [ - - - - ]
	//    [ - - - - ]
	//    [ - - - - ]
	{
		x, y := 0, VP-1
		deps[x][y].PostSet[H] = chs.Finish
	}

	// Horizontal dependecy
	//    [ H H H - ]
	//    [ - H H - ]
	//    [ - - H - ]
	//    [ - - - - ]
	for x := 0; x < VP-1; x++ {
		for y := 0; y < VP-1; y++ {
			// Tree[x][y] -> relay -> Tree[x][y+1]
			relay := chs.H[x][y]

			// Tree[x][y] -> H[x][y]
			deps[x][y].PostSet[H] = relay
			// H[x][y] -> Tree[x][y+1]
			deps[x][y+1].PreSet[H] = relay
		}
	}

	// Vertical dependecy
	//    [ - - - - ]
	//    [ - V V V ]
	//    [ - - V V ]
	//    [ - - - V ]
	for x := 1; x < VP; x++ {
		for y := 1; y < VP; y++ {
			// Tree[x][y] -> relay -> Tree[x-1][y]
			relay := chs.V[x-1][y-1]

			// Tree[x][y] -> V[x-1][y-1]
			deps[x][y].PostSet[V] = relay
			// V[x-1][y-1] -> Tree[x-1][y]
			deps[x-1][y].PreSet[V] = relay
		}
	}

	// width of each chunk
	width := N / VP

	// Delegate chunks to goroutine
	for x, i := VP-1, N-width; x >= 0; x-- {
		for y, j := x, i+1; y < VP; y++ {

			// Make goroutine to process chunk.
			// Each chunk is aware of dependency
			go Chunk(tree, i, j, width, &deps[x][y])

			j += width // jump column
		}
		i -= width // jump row
	}

	// Wait for the last goroutine.
	//  - Makes the dependency matrix alive until the end
	// See Also : Chunk()
	<-chs.Finish

	return
}
