package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"watch"
)

const (
	// Problem Size
	// Maximum N for machine is (1 << 13)
	N = 1 << 11

	// Tile Size
	VP = 1 << 11
)

var (
	h, v   [VP][VP]chan int
	finish chan int

	cost [N + 1][N + 1]float64
	root [N + 1][N + 1]int
	prob [N]float64
)

func main() {

	// Number of Physical Processer
	NP := runtime.NumCPU()
	runtime.GOMAXPROCS(NP)

	fmt.Printf("[Proc] : %d \n", NP)
	fmt.Printf("[N]    : %d \n", N)
	fmt.Printf("[VP]   : %d \n", VP)

	Setup()
	w := watch.StopWatch{}

	// Calculation
	// ---- ---- ---- ---- ----
	w.Reset()
	for d := 0; d < VP; d++ { //sub-diagonal of j=i+d
		for i := 0; i+d < VP; i++ {
			go Chunk(i, i+d)
		}
	}
	<-finish

	dur := w.Reset()
	// ---- ---- ---- ---- ----

	milisec := dur.Nanoseconds() / 1000000 // ms conversion
	fmt.Printf("[Time] : \t%d ms \n", milisec)
}

// Initialize variables
func Setup() {
	// ---- ---- ---- ---- ----
	var total float64
	for i := 0; i < N; i++ {
		var value float64 = rand.Float64()
		total += value
		prob[i] = value
	}

	for i := 0; i < N; i++ {
		prob[i] = prob[i] / total
	}
	// ---- ---- ---- ---- ----

	for i := 0; i < VP; i++ {
		for j := i; j < VP; j++ {
			if i < VP-1 {
				h[i][j] = make(chan int, 1)
			}
			if i > 0 {
				v[i][j] = make(chan int, 1)
			}
		}
	}

	finish = make(chan int, 1)
}

func Chunk(i, j int) {
	var bb int

	il := (i * (N + 1)) / VP //block-low for i
	jl := (j * (N + 1)) / VP //block-low for j

	ih := ((i+1)*(N+1))/VP - 1 //block-high for i
	jh := ((j+1)*(N+1))/VP - 1 //block-high for j

	if i < j { // not a tile on the diagonal
		<-h[i][j-1] // receive from the left
		<-v[i+1][j] // receive from below
	}

	for ii := ih; ii >= il; ii-- {
		if i == j {
			bb = ii
		} else {
			bb = jl
		}
		for jj := bb; jj <= jh; jj++ {
			root[ii][jj], cost[ii][jj] = Tree(ii, jj)
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

func Tree(row, col int) (rt int, cst float64) {
	var bestCost float64 = math.MaxFloat64
	var bestRoot int = -1

	switch {
	case row >= col: // Unused range
		rt, cst = -1, 0.0

	case row+1 == col: // Main diagonal
		rt, cst = row+1, prob[row]

	case row+1 < col: // Tree estimation...

		sum := 0.0 // basic cost
		for k := row; k < col; k++ {
			sum += prob[k]
		}

		// find optimized case
		for i := row; i < col; i++ {
			rCost := cost[row][i] + cost[i+1][col]
			if rCost < bestCost {
				bestCost = rCost
				bestRoot = i + 1
			}
		}

		rt, cst = bestRoot, bestCost+sum
	} //switch

	return
}
