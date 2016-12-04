package main

import (
	"flag"
	"fmt"
	"obst"
	"os"
	"research"
	"runtime"
	"watch"
)

var (
	N     int  // Problem's size
	NP    int  // Number of Physical Processer
	VP    int  // Chunk(Sub-problem) size
	isPar bool // Parallel execution flag

	cfg research.Config // Global Config
)

func init() {
	// Setup default values...
	N = 1 << 11
	NP = runtime.NumCPU() // Maximum core
	VP = NP * NP
	isPar = true

	flag.IntVar(&N, "n", N, "Problem's size")
	flag.IntVar(&NP, "np", NP, "Number of physical processor")
	flag.IntVar(&VP, "vp", VP, "Sub-problem's size")
	flag.BoolVar(&isPar, "parallel", isPar, "Parallel execution")
}

func setup() {
	if isPar == false {
		// Sequential execution
		NP = 1
	}
	runtime.GOMAXPROCS(NP)

	cfg.Init(N, NP, VP)
	cfg.Display(os.Stdout)
}

func main() {
	flag.Parse() // Parse the flags

	// ---- Construct / Initialize ----
	tree := obst.NewTree(N)
	tree.Init()

	setup() // Apply flags

	// ---- ---- Evaluation  ---- ----
	swatch := new(watch.StopWatch)
	swatch.Reset()
	if isPar == true {
		// Parallel
		research.EvaluatePar(&cfg, tree)
	} else {
		// Sequential
		research.EvaluateSeq(&cfg, tree)
	}

	// ---- ---- Result ---- ----
	duration := swatch.Pick()
	milisec := duration.Nanoseconds() / 1000000

	if isPar {
		fmt.Fprintf(os.Stdout, "[ %10s ] : %8d ms \n",
			"Parallel", milisec)
	} else {
		fmt.Fprintf(os.Stdout, "[ %10s ] : %8d ms \n",
			"Sequential", milisec)
	}

	return
}
