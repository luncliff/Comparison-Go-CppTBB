#ifndef _RESEARCH_OBST_H_
#define _RESEARCH_OBST_H_

#include <tuple>
#include <vector>

#include <tbb/task.h>

#include "./Alias.h"
#include "./utils/Matrix.h"


// OBST
struct Tree
{
    template <typename T>
    using Vector = std::vector<T>;

    Vector<f64>  prob;
    Matrix<f64>  cost{ prob.size() + 1, prob.size() + 1 };
    Matrix<i64>  root{ prob.size() + 1, prob.size() + 1 };

    explicit Tree(size_t _size) :
        prob(_size)
    {}

    static 
    auto Calculate(Tree& _tree, i32 _row, i32 _col) -> std::tuple<i64, f64>
    {
        i64  root{},   best_root = -1;
        f64  weight{}, best_weight = LDBL_MAX;

        // Unused range
        if (_row >= _col) {
            root = -1; weight = 0.0;
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
                                  + _tree.cost[i + 1][_col];

                if (temp_weight < best_weight) {
                    best_weight = temp_weight;
                    best_root = i + 1;
                }
            }

            root = best_root;
            weight = best_weight + sum;
        }

        return std::make_tuple(root, weight);
    }

    size_t size() const noexcept
    {
        return prob.size();
    }
};


static void Init(Tree& _tree)
{
    using namespace std;
    using namespace std::chrono;

    auto duration = system_clock::now().time_since_epoch();
    auto seed = static_cast<u32>(duration.count());
    mt19937 rnd{ seed };

    f64     sum{};
    for (i32 i = 0; i < _tree.prob.size(); ++i) {
        sum += _tree.prob[i] = rnd() % (1 << 20);
    }
    for (i32 i = 0; i < _tree.prob.size(); ++i) {
        _tree.prob[i] /= sum;
    }
}


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
            printf_s(" [%2d,%2.2lf] ", root, cost);
        }
        putchar('\n');
    }
}


#endif
