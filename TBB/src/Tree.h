// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : Tree.h
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2016/12/17
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

#include "./Alias.h"        // Alias for primitive types
#include "./utils/Matrix.h" // Custom Matrix for runtime 

template <typename T>
using Vector = std::vector<T>;

// - Note
//      Optimal Binary Search Tree for Dynamic Programming
struct Tree
{
    Vector<f64>  prob;  // Probability of vertices
    Matrix<f64>  cost;  // Cost matrix
    Matrix<i64>  root;  // Root index matrix

    // - Note
    //      Allocate memory for tree
    //      - prob[ N ]
    //      - cost[ N+1 ][ N+1 ]
    //      - root[ N+1 ][ N+1 ]
    explicit Tree(size_t n) :
        prob(n),
        cost{ n + 1 },
        root{ n + 1 }
    {}

    // - Note
    //      No side effect in the function. 
    //      Explicit assignment is required after calculation
    // - Example
    //      auto tuple = Calculate(tree, R, C);
    //      tree.root[R][C] = tuple.root;
    //      tree.cost[R][C] = tuple.cost;
    static 
    auto Calculate(const Tree& _tree, i32 _row, i32 _col) -> std::tuple<i64, f64>
    {
        i64  root{},   best_root = -1;
        f64  weight{}, best_weight = LDBL_MAX;

        // Unused range
        if (_row >= _col) {
            root = -1; 
            weight = 0.0;
        }
        // Main diagonal
        else if (_row + 1 == _col) {
            root = _row + 1;
            weight = _tree.prob[_row];
        }
        // Tree estimation
        else {
            // basic weight
            f64 sum = 0; 

            for (i32 i = _row; i < _col; ++i)
            {
                // Accumulate
                sum += _tree.prob[i];

                // Find best weight
                f64 temp_weight = _tree.cost[_row][i]
                                  + _tree.cost[ i+1 ][_col];

                if (temp_weight < best_weight) {
                    best_weight = temp_weight;
                    best_root = i + 1;
                }
            }
            root    = best_root;
            weight  = best_weight + sum;
        }
        // Weight == Cost
        return std::make_tuple(root, weight);
    }

    // - Note
    //      Number of vertices in tree
    size_t size() const noexcept
    {
        return prob.size();
    }
};

// - Note
//      Initialize tree's probabilities with random value
static void Init(Tree& _tree)
{
    using namespace std;
    using namespace std::chrono;

    // Use epoch time as the seed of random generator
    auto duration = system_clock::now().time_since_epoch();
    auto seed     = static_cast<u32>(duration.count());
    mt19937 rnd{ seed };

    f64     total{};
    for (i32 i = 0; i < _tree.prob.size(); ++i) {
        total += _tree.prob[i] = rnd() % (1 << 20);
    }
    // Normalization
    for (i32 i = 0; i < _tree.prob.size(); ++i) {
        _tree.prob[i] /= total;
    }
}

// - Note
//      Display the tree in console. 
//      Used for debugging
static void Display(const Tree& _tree)
{
    using namespace std;

    const i32 N = static_cast<i32>(_tree.size());

    puts("--- Prob ---");
    for (i32 i = 0; i < N; ++i) {
        printf_s(" %4.4lf, ", _tree.prob[i]);
    }
    putchar('\n');

    puts("--- Root/Cost ---");
    for (i32 i = 0; i < N + 1; ++i) {
        for (i32 j = 0; j < N + 1; ++j)
        {
            auto root = _tree.root[i][j];
            auto cost = _tree.cost[i][j];
            printf_s(" [%2lld,%2.2lf] ", root, cost);
        }
        putchar('\n');
    }
}


static 
bool operator==(const Tree& lhs, const Tree& rhs) noexcept
{
    if (lhs.prob != rhs.prob) {
        return false;
    }
    const i32 N = static_cast<i32>(lhs.size());

    // Compare valid range(x) only.
    //   [ - x x x ]
    //   [ - - x x ]
    //   [ - - - x ]
    //   [ - - - - ]
    for (i32 i = 0; i < N; ++i) {
        for (i32 j = i+1; j < N; ++j) 
        {
            const bool equal_root = lhs.root[i][j] != rhs.root[i][j];
            const bool equal_cost = lhs.cost[i][j] != rhs.cost[i][j];
            
            if ( equal_root == false || equal_cost == false){
                return false;
            }
        }
    }
    return true;
}

#endif
