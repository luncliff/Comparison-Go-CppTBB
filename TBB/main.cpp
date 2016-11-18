#include <iostream>
#include <thread>

#include <tbb\task.h>
#include "./DagTask.hpp"

using namespace std;
using namespace tbb;

template <typename T>
struct Matrix
{
    T** pp;
    

    Matrix(T** _pp = nullptr) : pp{_pp}
    {}

    operator T**() noexcept 
    {
        return pp;
    }
};

static double Evaluate(int vp) 
{
    int     n;
    double  *prob;

    //double  **cost;
    //int     **root;
    Matrix<double> cost;
    Matrix<int> root;


    // Matrix of DagTask*
    //DagTask*** x; 
    Matrix<DagTask*> x;

    // ...
    //allocate,initialize arrays cost,root,prob 
    // ... 

    //create tasks 
    for (int d = 0; d < vp; d++) {
        for (int i = 0; i + d < vp; i++)
        {
            // DagTask allocation with placement `new`
            x[i][i + d] = 
                new(tbb::task::allocate_root()) 
                DagTask{ i, i + d, vp, n, prob, cost, root };

            x[i][i + d]->set_ref_count(d == 0 ? 0 : 2);
        }
    }

    //set up successor links and front link 
    for (int d = 0; d < vp; d++){
        for (int i = 0; i + d < vp; i++)
        {
            // up 
            x[i][i + d]->successor[0] = i - 1 > -1 ? x[i - 1][i + d] : NULL;
            // right 
            x[i][i + d]->successor[1] = i + d + 1 < vp ? x[i][i + d + 1] : NULL;

            //main diagonal 
            if (d == 0 && i < vp - 1)
                x[i][i + d]->a = x[i + 1][i + 1];
        }
    }
                
    x[0][vp-1]->increment_ref_count(); 
    x[0][vp-1]->spawn_and_wait_for_all(*x[0][0]); 
    x[0][vp-1]->execute(); 
    tbb::task::destroy(*x[0][vp-1]); 
            
    return cost[0][n]; 
}


