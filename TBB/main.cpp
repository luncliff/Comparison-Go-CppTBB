#include <iostream>
#include <thread>

#include <tbb\task.h>
#include "ChunkTask.hpp"

using namespace std;
using namespace tbb;

struct Config
{
    int N, NP, VP;
};


// Sequential Evaluation
static void EvaliateSeq(const Config& cfg, Tree& _tree)
{
    int     N = cfg.N;

    for (int i = N-1; i >= 0; --i) {
        for (int j = 0; j < N; ++j) {
            // Calculate without side-effect
            auto root_cost = Tree::Calculate(_tree, i, j);

            // Assign result
            _tree.root[i][j] = std::get<0>(root_cost);
            _tree.cost[i][j] = std::get<1>(root_cost);
        }
    }

    return;
}

// Parallel Evaluation
static void EvaluatePar(const Config& cfg, Tree& _tree)
{
    int     vp  = cfg.VP;
    int     n   = cfg.N;
    // Matrix of ChunkTask.
    //  == Task[vp][vp]
    Matrix<tbb::task*> mtask{ vp, vp };

    // ---- ---- Construction ---- ---- ----

    for (int x = 0; x < vp - 1; ++x) {
        for (int y = x; y < vp - 1; ++y) 
        {
            auto place = tbb::task::allocate_root();
            // Correct index?
            mtask[x][y] = new(place) ChunkTask{ _tree, x, y, vp };
        }
    }

    // ---- ---- Setup : Relationship ---- ---- ----

    for (int x = 0; x < vp - 1; ++x) {
        for (int y = x; y < vp - 1; ++y)
        {
            x[i][i + d]->set_ref_count(d == 0 ? 0 : 2);

            x[i][i + d]->post_set[0] = i - 1 > -1 ? x[i - 1][i + d] : NULL;
            // right 
            x[i][i + d]->post_set[1] = i + d + 1 < vp ? x[i][i + d + 1] : NULL;

            //main diagonal 
            if (d == 0 && i < vp - 1)
                x[i][i + d]->a = x[i + 1][i + 1];

        }
    }

    // ---- ---- Trigger Execution ---- ---- ----

    x[0][vp - 1]->increment_ref_count();
    x[0][vp - 1]->spawn_and_wait_for_all(*x[0][0]);
    x[0][vp - 1]->execute();


    // ---- ---- Clean-Up/Return ---- ---- ----

    tbb::task::destroy(*x[0][vp-1]); 
            
    return;
}
