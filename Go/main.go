package main

import (
	"fmt"
)

const (
	n = 10
)

var (
	cost [n + 1][n + 1]float64
	root [n + 1][n + 1]int
	prob [n]float64
)

func mst(i, j int) {
	var bestCost float64 = 1e9 + 0.0
	var bestRoot int = -1

	switch {
	case i >= j:
		cost[i][j] = 0.0
		root[i][j] = -1
	case i+1 == j:
		cost[i][j] = prob[i]
		root[i][j] = i + 1
	case i+1 < j:
		psum := 0.0
		for k := i; k <= j-1; k++ {
			psum += prob[k]
		}
		for r := i; r <= j-1; r++ {
			rcost := cost[i][r] + cost[r+1][j]
			if rcost < bestCost {
				bestCost = rcost
				bestRoot = r + 1
			}
			cost[i][j] = bestCost + psum
			root[i][j] = bestRoot
		}
	} //switch

} // mst

func main() {
	// initialize prob[]
	for i := n; i >= 0; i-- {
		for j := i; j <= n; j++ {
			mst(i, j)
		}
	}
	for i := n; i >= 0; i-- {
		for j := i; j <= n; j++ {
			fmt.Println(root[i][j])
		}
	}

	fmt.Println("Hello, Go!")

}
