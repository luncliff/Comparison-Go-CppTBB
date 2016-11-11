package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"watch"
)

const (
	// Maximum N for machine is (1 << 13)
	N = 1 << 11

	Vp = 1 << 10
)

var (
	h, v   [Vp][Vp]chan int
	finish chan int

	cost [N + 1][N + 1]float64
	root [N + 1][N + 1]int
	prob [N]float64
)

// Initialize variables
func Init() {
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

	for i := 0; i < Vp; i++ {
		for j := i; j < Vp; j++ {
			if i < Vp-1 {
				h[i][j] = make(chan int, 1)
			}
			if i > 0 {
				v[i][j] = make(chan int, 1)
			}
		}
	}

	finish = make(chan int, 1)
}

func main() {
	fmt.Printf("[CPU] : %d \n", runtime.NumCPU())
	fmt.Printf("[N] : %d \n", N)
	fmt.Printf("[VP] : %d \n", Vp)

	// Maximize Concurrency
	runtime.GOMAXPROCS(runtime.NumCPU())

	w := watch.StopWatch{}

	Init()

	fmt.Println("Initialized")
	w.Reset()

	// Calculation
	// ---- ---- ---- ---- ----
	// for i := N; i >= 0; i-- {
	// 	for j := i; j <= N; j++ {
	// 		root[i][j], cost[i][j] = Sequential(i, j)
	// 	}
	// }
	// ---- ---- ---- ---- ----
	for d := 0; d < Vp; d++ { //sub-diagonal of j=i+d
		for i := 0; i+d < Vp; i++ {
			go Chunk(i, i+d)
		}
	}
	<-finish

	dur := w.Reset()
	sec := dur.Seconds()
	fmt.Printf("[Time] : %0.3f sec \n", sec)
	// nano := dur.Nanoseconds()
	// fmt.Printf("[Time] : %d ns \n", nano)

	// for i := 0; i < N; i++ { //sub-diagonal of j=i+d
	// 	for j := 0; j < N; j++ {
	// 		fmt.Printf("[%2d, %0.1f] ", root[i][j], cost[i][j])
	// 	}
	// 	fmt.Println()
	// }

}

func Chunk(i, j int) {
	var bb int

	il := (i * (N + 1)) / Vp //block-low for i
	jl := (j * (N + 1)) / Vp //block-low for j

	ih := ((i+1)*(N+1))/Vp - 1 //block-high for i
	jh := ((j+1)*(N+1))/Vp - 1 //block-high for j

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
			// mst(ii, jj)
			root[ii][jj], cost[ii][jj] = Sequential(ii, jj)
		}
	}

	if j < Vp-1 { // not a tile on the right border
		h[i][j] <- 1
	}
	if i > 0 { // not a tile on the top border
		v[i][j] <- 1
	}
	if i == 0 && j == Vp-1 { //the last tile
		finish <- 1
	}
}

func Sequential(r, c int) (Root int, Cost float64) {
	var bestCost float64 = math.MaxFloat64
	var bestRoot int = -1

	switch {
	case r >= c: // Unused range
		Root, Cost = -1, 0.0

	case r+1 == c: // Main diagonal
		Root, Cost = r+1, prob[r]

	case r+1 < c: // BST

		sum := 0.0 // basic cost
		for k := r; k < c; k++ {
			sum += prob[k]
		}

		// find optimized case
		for i := r; i < c; i++ {
			rCost := cost[r][i] + cost[i+1][c]
			if rCost < bestCost {
				bestCost = rCost
				bestRoot = i + 1
			}
		}
		Root, Cost = bestRoot, bestCost+sum
	} //switch

	return
}
