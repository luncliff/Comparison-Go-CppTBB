package main

import (
	"fmt"
	"io"
	"matrix"
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

// Channels ...
//  	Shared channels for synchronization
type Channels struct {
	H, V   [][]chan int // Matrix for sync
	Finish chan int     // Notify finish
}

// Init ...
//		Initialize set of shared data
func (shd *Channels) Init(width int) {
	// Allocate Horizontal/Vertical channels
	// Square matrix
	shd.H = matrix.ChanInt2D(width, width)
	shd.V = matrix.ChanInt2D(width, width)

	for i := 0; i < width; i++ {
		for j := i; j < width; j++ {
			// Bounded capacity : 1
			shd.H[i][j] = make(chan int, 1)
			shd.V[i][j] = make(chan int, 1)
		}
	}

	// Finish notifier channel
	shd.Finish = make(chan int, 1)
}
