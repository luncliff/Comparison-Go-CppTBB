// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author 	: Park  Dong Ha ( luncliff@gmail.com )
//
// 	Note	:
//		Channels for synchronization in parallel evaluation
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package main

// Channels ...
//  	Shared channels for synchronization
type Channels struct {
	H, V   [][]chan int // Matrix for sync
	Finish chan int     // Notify finish
}

// Init ...
//		Initialize set of shared data
func (chs *Channels) Init(width int) {
	// Allocate Horizontal/Vertical channels
	// Square matrix
	chs.H = ChanInt2D(width, width)
	chs.V = ChanInt2D(width, width)

	// Range :
	//      [ + + + + ]
	//      [   + + + ]
	//      [     + + ]
	//      [       + ]
	for i := 0; i < width; i++ {
		for j := i; j < width; j++ {
			// Bounded capacity : 1
			chs.H[i][j] = make(chan int, 1)
			chs.V[i][j] = make(chan int, 1)
		}
	}

	// Finish notifier channel
	chs.Finish = make(chan int, 1)
}

// ChanInt2D ...
// 		Matrix implementation with nested slice
func ChanInt2D(row int, col int) (mat [][]chan int) {
	size := row * col
	arr := make([]chan int, size)

	for i := 0; i < size; i += col {
		part := arr[i : i+col]
		mat = append(mat, part)
	}
	return
}
