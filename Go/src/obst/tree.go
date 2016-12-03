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
	cost [][]float64
	root [][]int
	prob []float64
}

// NewTree ...
//		Allocate memory resoureces for Optimal Binary Search Tree
//  	See also : type `Tree`
func NewTree(n uint) *Tree {
	res := new(Tree) // allocate
	res.prob = make([]float64, n)

	res.root = matrix.Int2D(n+1, n+1)     // `[N+1][N+1]int`
	res.cost = matrix.Float642D(n+1, n+1) // `[N+1][N+1]float64`

	return res
}

// Init ...
//  	Setup probabilities of the tree's vertices
//  	Receiver
//  		_tree
func (tree *Tree) Init() {
	var total float64
	// Distribute random values
	for i := range tree.prob {
		R := rand.Float64()
		total += R
		tree.prob[i] = R
	}
	// Normalize with total value
	for i := range tree.prob {
		tree.prob[i] = tree.prob[i] / total
	}
}

// Size ...
//  	The number of vertices in tree
func (tree *Tree) Size() uint {
	return len(tree.prob)
}

// Calculate ...
//  	To prevent side effect, explicit assignment is required
//  	after calculation. Tree won't be modified in the function
//  	- Receiver
//  		tree
//  	- Params
//  		row
//  		col
// 		- Returns
//  		root
//  		weight
func (tree *Tree) Calculate(row int, col int) (root int, weight float64) {

	var bestWeight float64 = math.MaxFloat64 // Optimal Cost
	var bestRoot int = -1                    // Optimal Root

	switch {
	// Unused range
	case row >= col:
		root, weight = -1, 0.0

	// Main diagonal
	case row+1 == col:
		root, weight = row+1, tree.prob[row]

	// Estimation : (Data Dependency Exists)
	case row+1 < col:
		sum := 0.0 // basic weight of tree

		for i := row; i < col; i++ {
			sum += tree.prob[i] // Accumulate

			// Find optimized case
			tempWeight := tree.cost[row][i] + tree.cost[i+1][col]
			if tempWeight < bestWeight {
				bestWeight = tempWeight
				bestRoot = i + 1 // 1-based indexing
			}
		}
		// Root : Optimal root index
		// Cost : Sum of all cost + Additinal weight
		root, weight = bestRoot, bestWeight+sum
	} //switch
	return
}
