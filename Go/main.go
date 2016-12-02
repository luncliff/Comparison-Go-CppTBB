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

// Config ...
//  	Test data
type Config struct {
	// N  : Problem's size
	// NP : Number of processors
	// VP : Chunk size
	N, NP, VP int
}

// Setup ...
// 		Initialize global variables
//  	Receiver
//  		_cfg : Config
func (_cfg *Config) Setup(n, np, vp int) {
	// Copy Constants
	_cfg.N, _cfg.NP, _cfg.VP = n, np, vp

}

// Display ...
//
func (_cfg *Config) Display(writer io.Writer) {
	fmt.Fprintf(writer, "[Proc] : %5d \n", _cfg.NP)
	fmt.Fprintf(writer, "[N]    : %5d \n", _cfg.N)
	fmt.Fprintf(writer, "[VP]   : %5d \n", _cfg.VP)
}

// Shared ...
//  	Shared data for synchronization
type Shared struct {
	// Channel matrix for sychronizatin
	h, v [][]chan int
	// Channel to notify finish
	finish chan int
}

// Setup ...
//		Initialize set of shared data
func (_shd *Shared) Setup(n, np, vp int) {
	// Allocate Horizontal/Vertical channels
	_shd.h = make([][]chan int, vp)
	_shd.v = make([][]chan int, vp)

	for i := range _shd.h {
		// Bounded
		_shd.h[i] = make([]chan int, 1)
		_shd.v[i] = make([]chan int, 1)
	}

	// Finish notifier channel
	_shd.finish = make(chan int, 1)
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
	cfg Config          // Global Config
	shd Shared          // Sync data
	sw  watch.StopWatch // Stop watch
)

func init() {
	// Explicit Setting for Number of Physical Processer
	MaxNP := runtime.NumCPU()
	runtime.GOMAXPROCS(MaxNP)

	// Maximize
	cfg.Setup(MaxN, MaxNP, MaxVP)
	cfg.Display(os.Stdout)

	shd.Setup(MaxN, MaxNP, MaxVP)

}

func main() {
	// OBST to calculate
	tree := obst.NewTree(MaxN)
	tree.Setup()

	sw.Reset() // Start the timer

	Evaluate(&cfg, &shd, tree) // Process the problem

	dur := sw.Pick() // Timer result
	milisec := dur.Nanoseconds() / 1000000
	fmt.Fprintf(os.Stdout, "[Time] : \t%d ms \n", milisec)
}

// Evaluate ...
//  	Parallel Chunk processing
func Evaluate(cfg *Config, shd *Shared, tree *obst.Tree) {
	shd.finish <- 1
	for d := 0; d < cfg.VP; d++ { //sub-diagonal of j=i+d
		for i := 0; i+d < cfg.VP; i++ {
			// go Chunk(i, i+d)
		}
	}
	<-shd.finish
	return
}
