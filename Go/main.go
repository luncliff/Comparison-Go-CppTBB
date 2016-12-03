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
	cfg.Init(MaxN, MaxNP, MaxVP)
	shd.Init(cfg.VP)

	cfg.Display(os.Stdout)
}

func main() {
	// OBST to calculate
	tree := obst.NewTree(MaxN)
	tree.Setup()

	sw.Reset() // Start the timer

	Evaluate(&cfg, &shd, tree) // Process the problem

	duration := sw.Pick() // Timer result
	milisec := duration.Nanoseconds() / 1000000
	fmt.Fprintf(os.Stdout, "[Time] : \t%d ms \n", milisec)
}
