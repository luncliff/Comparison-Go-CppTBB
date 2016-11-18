#pragma once
#include <cstdint>
#include <tuple>
#include <tbb/task.h>

double* prob;
double** cost;
int** root;

static constexpr auto MaxDouble = LDBL_MAX;

auto Tree(int row, int col) 
{
    int     rt;     // return root
    double  cst;    // return cost

    int     bestRoot = -1;
    double  bestCost = MaxDouble;

    // Unused range
    if (row >= col) {
        rt = -1;
        cst = 0.0;
    }
    // Main diagonal
    else if (row + 1 == col) {
        rt = row + 1;
        cst = prob[row];
    }
    // Tree estimation
    else {

        double sum = 0; // basic cost
        for (auto k = row; k < col; ++k) {
            sum += prob[k];
        }

        for (auto i = row; i < col; ++i) {
            auto rCost = cost[row][i] + cost[i + 1][col];
            if (rCost < bestCost) {
                bestCost = rCost;
                bestRoot = i + 1;
            }
        }

        rt = bestRoot;
        cst = bestCost + sum;
    }

    return std::make_tuple(rt , cst);
}

const int N;
const int VP;

void chunk(int i, int j) 
{
    int bb{};

    auto il = (i*(N + 1)) / VP; // block-low for i
    auto jl = (j*(N + 1)) / VP; // block-low for j

    auto ih = (((i + 1)*(N + 1)) / VP) - 1; // block-high for i
    auto jh = (((j + 1)*(N + 1)) / VP) - 1; // block-high for j

    // Not diagonal
    if (i < j) {
        // Receive from left(Horizontal)
        // Receive from below(Vertical)
    }
    
    // Calculation
    for (auto ii = ih; ii >= il; --ii) {
        if (i == j) {
            bb == ii;
        }
        else {
            bb = jl;
        }
        for (auto jj = bb; jj <= jh; ++jj) {
            auto tuple = Tree(ii, jj);
            root[ii][jj] = std::get<0>(tuple);
            cost[ii][jj] = std::get<1>(tuple);
        }
    }

    if (j < VP - 1) {
        // Notify to right
    }
    if (i > 0) {
        // notify to up
    }
    if (i == 0 && j == VP - 1) {
        // finish
    }
}






class DagTask :
    public tbb::task
{
    const int i, j, vp, n;
    double* const prob;
    double** const cost;
    int** const root;

public:
    DagTask* successor[2];
    DagTask* a; //next task of the front 

    DagTask(int i_, int j_, int vp_, int n_,
            double* prob_,
            double** cost_,
            int** root_) :
        i(i_), j(j_), vp(vp_), n(n_),
        prob(prob_),
        cost(cost_),
        root(root_)
    {}

    tbb::task* execute() override
    {
        if (i == j && i != vp - 1) { //main diagonal 
            spawn(*a); //front node 
        }

        chunk(i, j, vp, n, prob, cost, root);

        for (int k = 0; k<2; ++k)
            if (DagTask* t = successor[k])

        if (t->decrement_ref_count() == 0)
            spawn(*t);

        return NULL;
    }
};

