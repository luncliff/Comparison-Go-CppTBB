package obst

import (
	"math"
	"math/rand"
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
func NewTree(_size int) *Tree {
	res := new(Tree)
	res.prob = make([]float64, _size)

	// 2-Dimensional Slice instead of Matrix
	// Construct `[size][size]int`
	// Check for later optimization
	res.root = make([][]int, _size)
	for i := range res.root {
		res.root[i] = make([]int, _size)
	}

	// 2-Dimensional Slice instead of Matrix
	// Construct `[size][size]float64`
	// Check for later optimization
	res.cost = make([][]float64, _size)
	for i := range res.cost {
		res.cost[i] = make([]float64, _size)
	}

	return res
}

// Setup ...
//  	Setup probabilities of the tree's vertices
//  	Receiver
//  		_tree
func (_tree *Tree) Setup() {
	var total float64
	// Set random values
	for i := range _tree.prob {
		rndVal := rand.Float64()
		total += rndVal
		_tree.prob[i] = rndVal
	}

	// Normalize with total value
	for i := range _tree.prob {
		_tree.prob[i] = _tree.prob[i] / total
	}
}

// Calculate ...
//  	To prevent side effect, explicit assignment after calculation is required
//  	Receiver
//  		_tree
//  	Params
//  		_row
//  		_col
// 		Returns
//  		- root_
//  		- cost_
func (_tree *Tree) Calculate(_row int, _col int) (root int, cost float64) {
	// Optimal Cost
	var opCost float64 = math.MaxFloat64
	// Optimal Root
	var opRoot int = -1

	switch {
	// Unused range
	case _row >= _col:
		root, cost = -1, 0.0

	// Main diagonal
	case _row+1 == _col:
		root, cost = _row+1, _tree.prob[_row]

	// Estimation : (Data Dependency Exists)
	case _row+1 < _col:
		sumCost := 0.0 // basic cost
		for k := _row; k < _col; k++ {
			sumCost += _tree.prob[k]
		}

		// Find optimized case
		for i := _row; i < _col; i++ {
			rCost := _tree.cost[_row][i] + _tree.cost[i+1][_col]
			if rCost < opCost {
				opCost = rCost
				opRoot = i + 1
			}
		}

		// Root : Optimal Root
		// Cost : Sum of all cost + Optimal Additinal Cost
		root, cost = opRoot, opCost+sumCost
	} //switch
	return
}
