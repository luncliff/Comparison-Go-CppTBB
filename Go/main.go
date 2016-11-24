package main

import (
	"fmt"
	"io"
	"obst"
	"os"
	"runtime"
	"watch"
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Environment ...
//  	Test data
type Environment struct {
	// N  : Problem's size
	// NP : Number of processors
	// VP : Chunk size
	N, NP, VP int

	// Channel matrix for sychronizatin
	h, v [][]chan int
	// Channel to notify finish
	finish chan int
}

// Setup ...
// 		Initialize global variables
//  	Receiver
//  		_env : Environment
func (_env *Environment) Setup(n, np, vp int) {
	// Copy Constants
	_env.N, _env.NP, _env.VP = n, np, vp

	// Allocate Horizontal/Vertical channels
	_env.h = make([][]chan int, _env.VP)
	_env.v = make([][]chan int, _env.VP)

	for i := range _env.h {
		// Bounded
		_env.h[i] = make([]chan int, 1)
		_env.v[i] = make([]chan int, 1)
	}

	// Finish notifier channel
	_env.finish = make(chan int, 1)
}

// Display ...
//
func (_env *Environment) Display(writer io.Writer) {
	fmt.Fprintf(writer, "[Proc] : %5d \n", _env.NP)
	fmt.Fprintf(writer, "[N]    : %5d \n", _env.N)
	fmt.Fprintf(writer, "[VP]   : %5d \n", _env.VP)
}

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// ChunkData ...
//  	Data set for chunk calculation & synchronization
type ChunkData struct {
	PreSet  [2]chan int
	PostSet [2]chan int
}

// Wait ...
//  	Receiver
//  		ch
func (ch *ChunkData) Wait() {
	if ch.PreSet[0] != nil {
		<-ch.PreSet[0]
	}
	if ch.PreSet[1] != nil {
		<-ch.PreSet[1]
	}
	return
}

// Notify ...
// 		Receiver
//			ch
func (ch *ChunkData) Notify() {
	if ch.PostSet[0] != nil {
		ch.PostSet[0] <- 1
	}
	if ch.PostSet[1] != nil {
		ch.PostSet[1] <- 1
	}
	return
}

// CalculateChunk ...
//		Params
//  		_tree
//  		i
//  		j
//  		vp
//  	Returns
//  		Side Effect
/*
func CalculateChunk(_tree *obst.Tree, i int, j int, vp int) {
	var bb int
	il := (i * (N + 1)) / vp //block-low for i
	jl := (j * (N + 1)) / vp //block-low for j
	ih := ((i+1)*(N+1))/vp - 1 //block-high for i
	jh := ((j+1)*(N+1))/vp - 1 //block-high for j
	if i < j { // not a tile on the diagonal
		<-h[i][j-1] // receive from the left
		<-v[i+1][j] // receive from below
	}
	for ii := ih; ii >= il; ii-- {
		if i == j {
			bb = ii
		} else {
			bb = jl
		}
		for jj := bb; jj <= jh; jj++ {
			// Calculate ...
			root, cost := _tree.Calculate(ii, jj)
			// Assignment
			_tree.root[ii][jj] = root
			_tree.cost[ii][jj] = cost
		}
	}
	if j < VP-1 { // not a tile on the right border
		h[i][j] <- 1
	}
	if i > 0 { // not a tile on the top border
		v[i][j] <- 1
	}
	if i == 0 && j == VP-1 { //the last tile
		finish <- 1
	}
}
*/

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

const (
	// MaxN ...
	//  	Problem size
	// 		Maximum N is (1 << 13)
	MaxN = 1 << 11

	// MaxVP ...
	//  	Chunk's size
	MaxVP = 1 << 11
)

var (
	// Global environment
	env Environment
	sw  watch.StopWatch
)

func init() {
	// Explicit Setting for Number of Physical Processer
	MaxNP := runtime.NumCPU()
	runtime.GOMAXPROCS(MaxNP)

	// Maximize
	env.Setup(MaxN, MaxNP, MaxVP)
	env.Display(os.Stdout)

}

func main() {
	// OBST to calculate
	tree := obst.NewTree(MaxN)
	tree.Setup()

	sw.Reset() // Start the timer

	Evaluate(&env, tree) // Process the problem

	dur := sw.Pick() // Timer result
	milisec := dur.Nanoseconds() / 1000000
	fmt.Fprintf(os.Stdout, "[Time] : \t%d ms \n", milisec)
}

// Evaluate ...
//  	Parallel Chunk processing
func Evaluate(env *Environment, tree *obst.Tree) {
	env.finish <- 1
	for d := 0; d < env.VP; d++ { //sub-diagonal of j=i+d
		for i := 0; i+d < env.VP; i++ {
			// go Chunk(i, i+d)
		}
	}
	<-env.finish
	return
}
