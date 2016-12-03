package main

import (
	"fmt"
	"io"
)

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

// Shared ...
//  	Shared data for synchronization
type Shared struct {
	// Channel matrix for sychronizatin
	h, v [][]chan int
	// Channel to notify finish
	finish chan int
}

// Init ...
//		Initialize set of shared data
func (shd *Shared) Init(n uint) {
	// Allocate Horizontal/Vertical channels
	shd.h = make([][]chan int, n)
	shd.v = make([][]chan int, n)

	for i := range shd.h {
		// Bounded
		shd.h[i] = make([]chan int, 1)
		shd.v[i] = make([]chan int, 1)
	}

	// Finish notifier channel
	shd.finish = make(chan int, 1)
}
