// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note     :
//      Optimal Binary Search Tree for Dynamic Programming
//
//  See also :
//      `Matrix<T>`
//  Reference :
//      https://www.cs.auckland.ac.nz/software/AlgAnim/opt_bin.html
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _RESEARCH_OBST_H_
#define _RESEARCH_OBST_H_

#include <tuple>
#include <vector>
#include <chrono>
#include <random>

#include "./Alias.h"         // Alias for primitive types
#include "../utils/Matrix.h" // Custom Matrix for runtime

template <typename T>
using Vector = std::vector<T>;

// - Note
//      Optimal Binary Search Tree for Dynamic Programming
struct Tree
{
    Vector<f64> prob; // Probability of vertices
    Matrix<f64> cost; // Cost matrix
    Matrix<i64> root; // Root index matrix

    // - Note
    //      Allocate memory for tree
    //      - prob[ N ]
    //      - cost[ N+1 ][ N+1 ]
    //      - root[ N+1 ][ N+1 ]
    explicit Tree(size_t n) : prob(n),
                              cost{n + 1},
                              root{n + 1}
    {
    }

    // - Note
    //      No side effect in the function.
    //      Explicit assignment is required after calculation
    // - Example
    //      auto tuple = Calculate(tree, R, C);
    //      tree.root[R][C] = tuple.root;
    //      tree.cost[R][C] = tuple.cost;
    static auto Calculate(const Tree &tree,
                          i32 _row, i32 _col) noexcept
        -> std::tuple<i64, f64>;

    // - Note
    //      Number of vertices in tree
    size_t size() const noexcept;
};

bool operator==(const Tree &lhs, const Tree &rhs) noexcept;

// - Note
//      Initialize tree's probabilities with random value
void Init(Tree &_tree);

// - Note
//      Display the tree in console.
//      Used for debugging
void Display(const Tree &_tree);

#endif
