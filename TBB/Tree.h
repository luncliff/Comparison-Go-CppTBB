#ifndef _RESEARCH_OBST_H_
#define _RESEARCH_OBST_H_

#include <cstdint>
#include <tuple>
#include <tbb/task.h>
#include <vector>
#include "./Matrix.h"

// OBST
struct Tree
{
    template <typename T>
    using Vector = std::vector<T>;

    Vector<double>  prob;
    Matrix<double>  cost{ prob.size() + 1, prob.size() + 1 };
    Matrix<int>     root{ prob.size() + 1, prob.size() + 1 };

    explicit Tree(size_t _size) :
        prob(_size)
    {}

    static 
    auto Calculate(Tree& _tree, int _row, int _col) -> std::tuple<int, double>
    {
        int     root{},   best_root = -1;
        double  weight{}, best_weight = LDBL_MAX;

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

            double sum = 0; // basic weight
            for (auto k = _row; k < _col; ++k) {
                sum += _tree.prob[k];
            }

            for (auto i = _row; i < _col; ++i) 
            {
                auto temp_weight = _tree.cost[_row][i] 
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

};

#endif
