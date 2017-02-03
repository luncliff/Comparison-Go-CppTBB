package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"research"
	"runtime"
	"watch"
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

const (
	N  int = 1 << 11
	VP int = 1 << 11
	NP int = 8

	Parallel bool = true
)

var (
	Cost [N + 1][N + 1]float64
	Root [N + 1][N + 1]int
	Prob [N]float64
)

// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----

func init() {
	// Explicit Setting for Number of Physical Processor
	runtime.GOMAXPROCS(NP)

	// Initialize Probabilities

	var total float64
	// Distribute random values
	for i := range Prob {
		R := rand.Float64()
		total += R
		Prob[i] = R
	}
	// Normalize with total value
	for i := range Prob {
		Prob[i] = Prob[i] / total
	}
}

func main() {

	timer := new(watch.StopWatch)
	// Start the timer
	timer.Reset()

	if Parallel == true {
		createChan()
		ParallelEvaluate()
	} else {
		SequentialEvaluate()
	}

	dur := timer.Pick()
	milisec := dur.Nanoseconds() / 1000000

	{
		var rep research.Report
		rep.Config.N = N
		rep.Config.NP = NP
		rep.Config.VP = VP
		rep.Config.Parallel = Parallel
		rep.Elapsed = milisec
		fmt.Fprintf(os.Stdout, rep.ToJSON())
	}
}

func Calculate(row int, col int) (root int, weight float64) {
	var bestWeight float64 = math.MaxFloat64 // Optimal Cost
	var bestRoot int = -1                    // Optimal Root

	switch {
	// Unused range
	case row >= col:
		root, weight = -1, 0.0

	// Main diagonal
	case row+1 == col:
		root, weight = row+1, Prob[row]

	// Estimation : (Data Dependency Exists)
	case row+1 < col:
		sum := 0.0 // basic weight of tree

		for i := row; i < col; i++ {
			sum += Prob[i] // Accumulate

			// Find optimized case
			tempWeight := Cost[row][i] + Cost[i+1][col]
			if tempWeight < bestWeight {
				bestWeight = tempWeight
				bestRoot = i + 1 // 1-based indexing
			}
		}
		// Root : Optimal root index
		// Cost : Sum of all cost + Additinal weight
		root, weight = bestRoot, bestWeight+sum
	}

	return // return tuple
}

func SequentialEvaluate() {
	for i := N; i >= 0; i-- {
		for j := i; j <= N; j++ {
			// Calculate ...
			root, cost := Calculate(i, j)
			// Assignment
			Root[i][j] = root
			Cost[i][j] = cost
		}
	}

}

var (
	// Channel matrix for synchronization
	h, v [VP][VP]chan int
	// Channel to notify finish
	finish chan int
)

func createChan() {

	// Finish notifier channel
	finish = make(chan int, 1)

	for i := 0; i < VP; i++ {
		for j := i; j < VP; j++ {
			if j < VP-1 {
				h[i][j] = make(chan int, 1)
			}
			if i > 0 {
				v[i][j] = make(chan int, 1)
			}
		}
	}
}

func Chunk(i int, j int) {
	il := (i * (N + 1)) / VP   //block-low for i
	jl := (j * (N + 1)) / VP   //block-low for j
	ih := ((i+1)*(N+1))/VP - 1 //block-high for i
	jh := ((j+1)*(N+1))/VP - 1 //block-high for j

	if i < j { // not a tile on the diagonal
		<-h[i][j-1] // receive from the left
		<-v[i+1][j] // receive from below
	}

	var bb int
	for ii := ih; ii >= il; ii-- {

		if i == j {
			bb = ii
		} else {
			bb = jl
		}
		for jj := bb; jj <= jh; jj++ {
			// Calculate ...
			root, cost := Calculate(ii, jj)
			// Assignment
			Root[ii][jj] = root
			Cost[ii][jj] = cost
		}
	}
	if j < VP-1 { // not a tile on the right border
		h[i][j] <- 1
	}
	if i > 0 { // not a tile on the top border
		v[i][j] <- 1
	}
	if i == 0 && j == VP-1 { //the last tile
		finish <- 1
	}
}

func ParallelEvaluate() {
	for d := 0; d < VP; d++ { //sub-diagonal of j=i+d
		for i := 0; i+d < VP; i++ {
			go Chunk(i, i+d)
		}
	}
	<-finish
}

func Display() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("[%d:%0.2f]", Root[i][j], Cost[i][j])
		}
		fmt.Println()
	}
}
