package main

import (
	"fmt"
	"math"
	"math/rand"
	"watch"
)

const (
	// This is maximum N
	N = 1 << 13
)

var (
	cost [N + 1][N + 1]float64
	root [N + 1][N + 1]int
	prob [N]float64
)

// Initialize variables
func init() {
	// ---- ---- ---- ---- ----
	var total float64
	for i := 0; i < N; i++ {
		var value float64 = rand.Float64() / 10
		total += value
		prob[i] = value
	}

	for i := 0; i < N; i++ {
		prob[i] = prob[i] / total
	}
	// ---- ---- ---- ---- ----

}

func main() {
	w := watch.StopWatch{}
	w.Reset()

	init()

	fmt.Printf("[Size] : %d \n", N)
	w.Reset()

	// Calculation
	// ---- ---- ---- ---- ----
	for i := N; i >= 0; i-- {
		for j := i; j <= N; j++ {
			root[i][j], cost[i][j] = Sequential(i, j)
		}
	}
	// ---- ---- ---- ---- ----

	dur := w.Reset()
	sec := dur.Seconds()
	fmt.Printf("[Time] : %0.3f sec \n", sec)
	// nano := dur.Nanoseconds()
	// fmt.Printf("[Time] : %d ns \n", nano)

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

} // mst
