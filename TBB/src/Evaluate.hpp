#ifndef _RESEARCH_EVALUATE_HPP_
#define _RESEARCH_EVALUATE_HPP_

#include "./Alias.h"
#include "./ChunkTask.hpp"

namespace Research
{
    struct Config
    {
        i32 N, NP, VP;
    };


    static void Display(const Config& _cfg)
    {
        using namespace std;

        printf_s(" [ Proc ] : %5d\n", _cfg.NP);
        printf_s(" [ N    ] : %5d\n", _cfg.N);
        printf_s(" [ VP   ] : %5d\n", _cfg.VP);
    }


    // ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----


    // Sequential Evaluation
    static 
    void EvaluateSeq(const Config& cfg, Tree& _tree) noexcept
    {
        const i32 N = cfg.N;

        // loop : bottom-left >>> top-right
        //      [ + + + + ]
        //      [   + + + ]
        //      [     + + ]
        //   -> [       + ]
        for (i32 i = N; i >= 0; --i) {
            for (i32 j = i; j <= N; ++j)
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
    static 
    void EvaluatePar(const Config& cfg, Tree& _tree) noexcept
    {
        // Delimit number of thread
        tbb::task_scheduler_init sched_init(cfg.NP);

        const i32  VP = cfg.VP;
        const i32  N = cfg.N;

        // Matrix of ChunkTask.
        //  == Task[vp][vp]
        Matrix<ChunkTask*> task{ static_cast<u32>(VP) };

        // ---- ---- Construction ---- ---- ----

        const i32 width = N / VP;
        // i, j : Tree matrices' index
        // x, y : Chunks' index
        for (i32 x = 0, i = 0; x < VP; ++x) {
            for (i32 y = x, j = i + 1; y < VP; ++y)
            {
                auto place = tbb::task::allocate_root();
                // Construct task to process task-problem
                task[x][y] = new(place) ChunkTask{ _tree, i, j, N / VP };

                // Most tasks' ref-counts are 2
                task[x][y]->set_ref_count(2);

                j += width;
            }
            i += width;
        }

        // ---- ---- Setup : Relationship ---- ---- ----
        const i32 begin = 0;
        const i32 end = VP - 1;

        for (i32 x = begin; x <= end; ++x) {

            // Main diagonal tasks have no dependency
            // make chained for fast spawning
            task[x][x]->set_ref_count(0);
            if (x < end) {
                task[x][x]->chain = task[x + 1][x + 1];
            }

            for (i32 y = x; y <= end; ++y)
            {
                static constexpr auto V = 0, H = 1;

                // Vertical/Horizontal successor tasks
                tbb::task* ver = (x == begin) ? nullptr : task[x - 1][y];
                tbb::task* hor = (y == end) ? nullptr : task[x][y + 1];

                task[x][y]->post_set[V] = ver;
                task[x][y]->post_set[H] = hor;
            }
        }


        // ---- ---- Trigger Execution ---- ---- ----

        // The last task
        tbb::task& last_task = *task[begin][end];
        if (VP > 1) {
            last_task.increment_ref_count();
            last_task.spawn_and_wait_for_all(*task[0][0]);
        }
        last_task.execute();

        // ---- ---- Clean-Up/Return ---- ---- ----

        tbb::task::destroy(last_task);

        return;
    }

}
#endif
