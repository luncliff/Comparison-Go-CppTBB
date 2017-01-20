// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File	: config.go
//  Author	: Park Dong Ha ( luncliff@gmail.com )
//  Updated	: 2016/12/17
//
// 	Note	:
//		Experiment configuration
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package research

import (
	"flag"
	"fmt"
	"io"
	"runtime"
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Config ...
//  	Test configuration
type Config struct {
	// N  : Problem's size
	// NP : Number of processors
	// VP : Chunk size
	N, NP, VP int
}

// Display ...
//      Display configuration via `io.Writer`
func (cfg *Config) Display(writer io.Writer) {
	fmt.Fprintf(writer, "[ Proc ] : %5d \n", cfg.NP)
	fmt.Fprintf(writer, "[ N    ] : %5d \n", cfg.N)
	fmt.Fprintf(writer, "[ VP   ] : %5d \n", cfg.VP)
}

// Parser ...
//  	Custom command line flag parser for this research
type Parser struct {
	N   int  // Problem's size
	NP  int  // Number of Physical Processer
	VP  int  // Scale of Sub-problems
	Par bool // Parallel execution flag
}

func (p *Parser) Init() {

	// Setup default values...
	p.N = 1 << 11           // 2048
	p.NP = runtime.NumCPU() // Maximum core
	p.VP = p.NP * p.NP      // Square of NP

	flag.IntVar(&p.N, "n", p.N, "Problem's size")
	flag.IntVar(&p.NP, "np", p.NP, "Number of physical processor")
	flag.IntVar(&p.VP, "vp", p.VP, "Sub-problem's size")
	flag.BoolVar(&p.Par, "parallel", p.Par, "Parallel execution")
}

func (p *Parser) Parse() {
	flag.Parse() // Parse the flags
}

func (p *Parser) Config() (cfg Config) {
	cfg.N = p.N
	if p.Par == false {
		// Sequential execution
		cfg.NP = 1
		cfg.VP = 1
	} else {
		cfg.NP = p.NP
		cfg.VP = p.VP
	}
	return
}
