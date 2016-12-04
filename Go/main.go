package main

import (
	"fmt"
	"obst"
	"os"
	"runtime"
	"watch"
)

const (
	// MaxN ...
	//  	Problem size
	// 		Maximum N is (1 << 13)
	MaxN = 1 << 2

	// MaxVP ...
	//  	Chunk's size
	MaxVP = 1 << 1
)

var (
	cfg Config          // Global Config
	chs Channels        // Shared Channels
	sw  watch.StopWatch // Stop watch
)

func init() {
	// Explicit Setting for Number of Physical Processer
	MaxNP := runtime.NumCPU()
	runtime.GOMAXPROCS(MaxNP)

	// Maximize
	cfg.Init(MaxN, MaxNP, MaxVP)
	chs.Init(cfg.VP - 1)

	cfg.Display(os.Stdout)
}

func main() {
	// OBST to calculate
	tree := obst.NewTree(MaxN)
	tree.Init()

	sw.Reset() // Start the timer

	EvaluatePar(&cfg, tree, &chs) // Process the problem

	duration := sw.Pick() // Timer result
	milisec := duration.Nanoseconds() / 1000000
	fmt.Fprintf(os.Stdout, "[Time] : \t%d ms \n", milisec)
}
