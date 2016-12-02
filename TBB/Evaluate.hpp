#pragma once
#include "./Config.h"
#include "./ChunkTask.hpp"

static void EvaliateSeq(const Config& cfg, Tree& _tree) noexcept;
static void EvaluatePar(const Config& cfg, Tree& _tree) noexcept;


// Sequential Evaluation
static void EvaliateSeq(const Config& cfg, Tree& _tree) noexcept
{
    const int N = cfg.N;

    // loop : bottom-left >>> top-right
    for (int i = N - 1; i >= 0; --i) {
        for (int j = i; j < N; ++j)
        {
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
//  - Reference
//      https://software.intel.com/en-us/node/506110
//  - Note
//      Chained spawning, Explicit task destroy
static void EvaluatePar(const Config& cfg, Tree& _tree) noexcept
{
    const int  VP = cfg.VP;
    const int  N  = cfg.N;

    // Matrix of ChunkTask.
    //  == Task[vp][vp]
    Matrix<ChunkTask*> task{ VP, VP };

    // ---- ---- Construction ---- ---- ----

    // i, j : tree's cell
    // x, y : chunk
    for (int x = 0, i = 0; i < N; ) {
        for (int y = x, j = i; j < N;) 
        {
            auto place = tbb::task::allocate_root();
            // Construct task to process task-problem
            task[x][y] = new(place) ChunkTask{ _tree, i, j, VP };

            // Most tasks' ref-counts are 2
            task[x][y]->set_ref_count(2);

            j += VP;
            ++y;
        }
        i += VP;
        ++x;
    }

    // ---- ---- Setup : Relationship ---- ---- ----
    const auto begin = 0;
    const auto end =VP - 1;

    for (int x = begin; x < end; ++x) {
        for (int y = x; y < end; ++y)
        {
            static constexpr auto V = 0, H = 1;

            // Main diagonal tasks have no dependency
            // make chained for fast spawning
            if (x == y) {   
                task[x][y]->set_ref_count(0);  
                task[x][y]->chain = task[ x+1 ][ y+1 ];
            }

            // Vertical/Horizontal successor tasks
            tbb::task* ver = (x == begin) ? nullptr : task[ x-1 ][ y ];
            tbb::task* hor = (y == end)   ? nullptr : task[ x ][ y+1 ];

            task[x][y]->post_set[V] = ver;
            task[x][y]->post_set[H] = hor;
        }
    }

    // ---- ---- Trigger Execution ---- ---- ----

    // The last task
    tbb::task& last_task = *task[begin][end];
    last_task.increment_ref_count();
    last_task.spawn_and_wait_for_all( *task[0][0] );
    last_task.execute();

    // ---- ---- Clean-Up/Return ---- ---- ----

    tbb::task::destroy(last_task);

    return;
}

