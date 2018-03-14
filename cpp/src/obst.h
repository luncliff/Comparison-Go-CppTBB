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

#include "heap_matrix.hpp"
// - Note
//      Optimal Binary Search Tree for Dynamic Programming
struct obst
{
    std::vector<float> prob; // Probability of vertices
    heap_matrix<int32_t> root;    // Root index matrix
    heap_matrix<float> cost;    // Cost matrix

public:
    // - Note
    //      Allocate memory for tree
    //      - prob[ N ]
    //      - cost[ N+1 ][ N+1 ]
    //      - root[ N+1 ][ N+1 ]
    explicit obst(std::size_t N) noexcept(false);
    ~obst() noexcept = default;

    // - Note
    //      Number of vertices in tree
    size_t size() const noexcept;

    void randomize() noexcept;
public:

    // - Note
    //      No side effect in the function.
    //      Explicit assignment is required after calculation
    // - Example
    //      auto tuple = Calculate(tree, R, C);
    //      tree.root[R][C] = tuple.root;
    //      tree.cost[R][C] = tuple.cost;
    static 
    auto optimal(const obst& tree,
                   uint32_t _row, uint32_t _col) noexcept
        -> std::tuple<int32_t, float>;

};

bool operator==(const obst& lhs, const obst& rhs) noexcept;
bool operator!=(const obst& lhs, const obst& rhs) noexcept;

//
//// - Note
////      Display the tree in console.
////      Used for debugging
//void Display(const Tree &_tree);

#endif
