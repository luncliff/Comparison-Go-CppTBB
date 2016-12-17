// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File	: tree.go
//  Author	: Park Dong Ha ( luncliff@gmail.com )
//  Updated	: 2016/12/17
//
// 	Note	:
//      Optimal Binary Search Tree for Dynamic Programming
//
//  See also :
//      `package matrix`
//  Reference :
//      https://www.cs.auckland.ac.nz/software/AlgAnim/opt_bin.html
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
package obst

import (
	"math"
	"math/rand"
	"matrix"
)

// Tree ...
//  	Custom structure for Optimal Binary Search Tree
//  	- `prob` : []float64
// 		- `root` : [][]int
//		- `cost` : [][]float64
type Tree struct {
	Prob []float64
	Cost [][]float64
	Root [][]int
}

// NewTree ...
//		Allocate memory resoureces for Optimal Binary Search Tree
//  	See also : type `Tree`
func NewTree(n int) *Tree {
	res := new(Tree) // allocate
	res.Prob = make([]float64, n)

	res.Root = matrix.Int2D(n+1, n+1)     // `[N+1][N+1]int`
	res.Cost = matrix.Float642D(n+1, n+1) // `[N+1][N+1]float64`

	return res
}

// Init ...
//  	Setup probabilities of the tree's vertices
//  	Receiver
//  		_tree
func (tree *Tree) Init() {
	var total float64
	// Distribute random values
	for i := range tree.Prob {
		R := rand.Float64()
		total += R
		tree.Prob[i] = R
	}
	// Normalize with total value
	for i := range tree.Prob {
		tree.Prob[i] = tree.Prob[i] / total
	}
}

// Size ...
//  	The number of vertices in tree
func (tree *Tree) Size() int {
	return len(tree.Prob)
}

// Calculate ...
//  	To prevent side effect,
//		explicit assignment is required after calculation.
// 		Tree won't be modified in the function
func (tree *Tree) Calculate(row int, col int) (root int, weight float64) {
	var bestWeight float64 = math.MaxFloat64 // Optimal Cost
	var bestRoot int = -1                    // Optimal Root

	switch {
	// Unused range
	case row >= col:
		root, weight = -1, 0.0

	// Main diagonal
	case row+1 == col:
		root, weight = row+1, tree.Prob[row]

	// Estimation : (Data Dependency Exists)
	case row+1 < col:
		sum := 0.0 // basic weight of tree

		for i := row; i < col; i++ {
			sum += tree.Prob[i] // Accumulate

			// Find optimized case
			tempWeight := tree.Cost[row][i] + tree.Cost[i+1][col]
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
