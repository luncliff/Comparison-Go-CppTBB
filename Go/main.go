// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : main.go
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2016/12/17
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
	"fmt"
	"obst"
	"research"
	"runtime"
	"watch"
)

var (
	parser research.Parser
)

func init() {
	parser.Init()
}

func main() {
	// Parse the command flags
	parser.Parse()

	// ==== ==== Setup configuration  ==== ====

	config := parser.Config()

	// Delimit the number of threads
	runtime.GOMAXPROCS(config.NP)

	// ==== ==== Construct / Initialize ==== ====

	tree := obst.NewTree(config.N)
	tree.Init()

	// ==== ==== ==== Evaluation ==== ==== ====

	timer := new(watch.StopWatch)
	timer.Reset()

	if config.Parallel == true {
		// Processing + Blocking Garbage Collection
		research.EvaluatePar(tree, config.VP)
		runtime.GC()
	} else {
		// Processing
		research.EvaluateSeq(tree)
	}

	// ==== ==== ==== Result ==== ==== ====

	duration := timer.Pick()
	elapsed := duration.Nanoseconds() / 1000000

	{
		var rep research.Report
		rep.Config = config
		rep.Elapsed = elapsed

		fmt.Println(rep.ToJSON())
	}

	// tree.Display(os.Stdout)
	return
}
