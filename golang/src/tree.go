// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author 	: Park  Dong Ha ( luncliff@gmail.com )
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
package main

import (
	"fmt"
	"io"
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
	Cost matrix.Float2D // Float64 2D matrix
	Root matrix.Int2D   // Int 2D matrix
}

// NewTree ...
//		Allocate memory resources for Optimal Binary Search Tree
//  	See also : type `Tree`
func NewTree(n int) *Tree {
	res := new(Tree)
	res.Prob = make([]float64, n)
	res.Root = matrix.MakeInt2D(n+1, n+1)
	res.Cost = matrix.MakeFloat2D(n+1, n+1)

	return res
}

// Init ...
//  	Setup probabilities of the tree's vertices
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
			tempWeight := *tree.Cost.At(row, i) + *tree.Cost.At(i+1, col)
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

func (lhs *Tree) Copy(rhs *Tree) {
	N := rhs.Size()

	for i := 0; i < N; i++ {
		lhs.Prob[i] = rhs.Prob[i]
	}
	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			*lhs.Root.At(i, j) = *rhs.Root.At(i, j)
			*lhs.Cost.At(i, j) = *rhs.Cost.At(i, j)
		}
	}
}

// Equal...
//  	Check if both tree is equal. Only consider valid indexes
func (lhs *Tree) Equal(rhs *Tree) bool {
	// Equal Size?
	if lhs.Size() != rhs.Size() {
		return false
	}
	N := lhs.Size()

	// Equal Probability ?
	for i := 0; i < N; i++ {
		if lhs.Prob[i] != rhs.Prob[i] {
			// log.Println("Different Prob")
			return false
		}
	}

	// Equal Root & Cost ?
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			equalroot := *lhs.Root.At(i, j) != *rhs.Root.At(i, j)
			equalcost := *lhs.Cost.At(i, j) != *rhs.Cost.At(i, j)
			if equalroot == false || equalcost == false {
				return false
			}
		}
	}
	return true
}

func (tree *Tree) Display(out io.Writer) {
	N := tree.Size()

	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			fmt.Fprintf(out, " [%2d, %2.2f]",
				*tree.Root.At(i, j), *tree.Cost.At(i, j))
		}
		fmt.Fprintln(out)
	}
}
