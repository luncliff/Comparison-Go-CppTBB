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
	"encoding/json"
	"flag"
	"runtime"
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

// Config ...
//  	Test configuration
type Config struct {
	// N  : Problem's size
	// NP : Number of processors
	// VP : Chunk size
	N        int
	NP       int `json:"Proc"`
	VP       int
	Parallel bool
}

// ToJSON ...
//      Display configuration via `io.Writer`
func (cfg *Config) ToJSON() string {
	byt, _ := json.Marshal(cfg)
	return string(byt)
}

// Parser ...
//  	Custom command line flag parser for this research
type Parser struct {
	n        int  // Problem's size
	np       int  // Number of Physical Processer
	vp       int  // Scale of Sub-problems
	parallel bool // Parallel execution flag
}

// Init ...
//  	Initialize the parser
func (p *Parser) Init() {

	// Setup default values...
	p.n = 1 << 11           // 2048
	p.np = runtime.NumCPU() // Maximum core
	p.vp = p.np * p.np      // Square of NP
	p.parallel = true

	flag.IntVar(&p.n, "n", p.n, "Problem's size")
	flag.IntVar(&p.np, "np", p.np, "Number of physical processor")
	flag.IntVar(&p.vp, "vp", p.vp, "Sub-problem's size")
	flag.BoolVar(&p.parallel, "parallel", p.parallel, "Parallel execution")
}

// Parse ...
//  	Parse the command argument with flag package
func (p *Parser) Parse() {
	flag.Parse() // Parse the flags
}

// Config ...
//  	Create configuration from the parser's state
func (p *Parser) Config() (cfg Config) {
	cfg.N = p.n
	cfg.Parallel = p.parallel

	if p.parallel == false {
		// Sequential execution
		cfg.NP = 1
		cfg.VP = 1
	} else {
		cfg.NP = p.np
		cfg.VP = p.vp
	}
	return
}

// Report ...
//  	Configuration + Elapsed time
type Report struct {
	Config  Config `json:"Config"`
	Elapsed int64  `json:"Elapsed"`
}

// ToJSON ...
//      Display configuration via `io.Writer`
func (rep *Report) ToJSON() string {
	byt, _ := json.Marshal(rep)
	return string(byt)
}
