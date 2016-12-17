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
	"fmt"
	"io"
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

// Init ...
// 		Initialize configuration variables
//  	- Receiver
//  		cfg : Config
func (cfg *Config) Init(n, np, vp int) {
	// Copy Constants
	cfg.N, cfg.NP, cfg.VP = n, np, vp

}

// Display ...
//      Display configuration via `io.Writer`
func (cfg *Config) Display(writer io.Writer) {
	fmt.Fprintf(writer, "[ Proc ] : %5d \n", cfg.NP)
	fmt.Fprintf(writer, "[ N    ] : %5d \n", cfg.N)
	fmt.Fprintf(writer, "[ VP   ] : %5d \n", cfg.VP)
}
