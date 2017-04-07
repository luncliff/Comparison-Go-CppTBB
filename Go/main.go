// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note     :
//      Evaluate Optimal Binary Search Tree problem based on the
//      command-line options.
//      - `N`  : Problem size
//      - `NP` : Number of Processors
//      - `VP` : Scale of Sub-problems
//              Small : big  sub-problem, but low  sync cost
//              Big   : tiny sub-problem, but high sync cost
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package main

import (
	"flag"
	"fmt"
	"obst"
	_ "os"
	"research"
	"runtime"
	_ "runtime/pprof"
	"watch"
)

var (
	// Parser for configuration
	parser research.Parser

	prof_cpu string
	prof_mem string
)

func init() {
	flag.StringVar(&prof_cpu,
		"cpuprofile", "cpu.prof", "CPU profile to file")
	flag.StringVar(&prof_mem,
		"memprofile", "mem.mprof", "Memory profile to this file")
	parser.Init()
}

func main() {
	// Parse the command flags
	parser.Parse()

	// ==== ==== ==== Setup configuration  ==== ==== ==== ==== ====

	config := parser.Config()
	// cpu, _ := os.Create(prof_cpu) // CPU report
	// mem, _ := os.Create(prof_mem) // Memory report

	// ==== ==== ==== Construct / Initialize ==== ==== ==== ====

	tree := obst.NewTree(config.N)
	tree.Init()

	// Delimit the number of threads
	runtime.GOMAXPROCS(config.NP)
	// pprof.StartCPUProfile(cpu)

	// ==== ==== ==== Evaluation ==== ==== ==== ==== ==== ==== ====

	timer := new(watch.StopWatch)
	timer.Reset()

	// Processing + Blocking Garbage Collection
	if config.Parallel == true {
		research.EvaluatePar(tree, config.VP)
		// pprof.WriteHeapProfile(mem) // Parallel Processing

		runtime.GC()
		// pprof.WriteHeapProfile(mem) // After GC
	} else {
		// Processing
		research.EvaluateSeq(tree)
	}

	// ==== ==== ==== Result ==== ==== ==== ==== ==== ==== ====

	elapsed := timer.Pick().Nanoseconds() / 1000000
	// pprof.StopCPUProfile() // End : CPU profile
	// mem.Close()            // End : Memory profile
	{
		var rep research.Report
		rep.Config = config
		rep.Elapsed = elapsed

		fmt.Println(rep.ToJSON())
	}

	// tree.Display(os.Stdout)
	return
}
