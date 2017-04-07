#include "./Evaluate.h"

namespace Research
{
// - Note
//      Sequential Evaluation
void EvaluateSeq(Tree &tree) noexcept
{
    const i32 N = static_cast<i32>(tree.size());

    // loop : bottom-left >>> top-right
    //      [ + + + + ]
    //      [   + + + ]
    //      [     + + ]
    //   -> [       + ]
    for (i32 i = N; i >= 0; --i)
    {
        for (i32 j = i; j <= N; ++j)
        {
            std::tie(tree.root[i][j], tree.cost[i][j]) 
            = Tree::Calculate(tree, i, j);

            // // Calculate without side-effect
            // auto root_cost = Tree::Calculate(tree, i, j);
            // // Assign result
            // tree.root[i][j] = std::get<0>(root_cost);
            // tree.cost[i][j] = std::get<1>(root_cost);
        }
    }
    return;
}

// - Note
//      Parallel Evaluation
//      Chained spawning, Explicit task destroy
// - Reference
//      https://software.intel.com/en-us/node/506110
void EvaluatePar(Tree &tree, const i32 VP) noexcept
{
    const i32 N = static_cast<i32>(tree.size());

    // Matrix of ChunkTask.
    //  == Task[vp][vp]
    Matrix<ChunkTask *> task{static_cast<u32>(VP)};

    // ---- ---- Construction ---- ---- ----

    const i32 width = N / VP;
    // i, j : Tree matrices' index
    // x, y : Chunks' index
    // - Range :
    //      [ + + + + ]
    //      [   + + + ]
    //      [     + + ]
    //      [       + ]
    for (i32 x = 0, i = 0; x < VP; ++x)
    {
        for (i32 y = x, j = i + 1; y < VP; ++y)
        {
            auto place = tbb::task::allocate_root();
            // Construct task to process task-problem
            task[x][y] = new (place) ChunkTask{tree, i, j, N / VP};

            // Most of ref-counts are 2
            task[x][y]->set_ref_count(2);

            j += width; // jump column
        }
        i += width; // jump row
    }

    // ---- ---- Setup : Relationship ---- ---- ----

    const i32 begin = 0;
    const i32 end = VP - 1;

    for (i32 x = begin; x <= end; ++x)
    {
        // Main diagonal tasks have no dependency
        task[x][x]->set_ref_count(0);

        // Make chained for fast spawning
        //   -> [ 1       ]     // 1 spawns 2...
        //      [   2     ]     // 2 spawns 3...
        //      [     3   ]     // 3 spawns 4...
        //      [       4 ]     // no spawning
        // Exclude the last chunk at main diagonal
        if (x < end)
        {
            task[x][x]->chain = task[x + 1][x + 1];
        }

        // - Dependency
        //                 t[x-1][y]
        //                     ^
        //    t[x][y-1]  >  t[x][y]   > t[x][y+1]
        //                     ^
        //                 t[x+1][y]
        for (i32 y = x; y <= end; ++y)
        {
            static constexpr auto V = 0, H = 1;
            // Vertical/Horizontal successor tasks(post set)
            tbb::task *ver = (x == begin) ? nullptr : task[x - 1][y];
            tbb::task *hor = (y == end) ? nullptr : task[x][y + 1];

            task[x][y]->post_set[V] = ver;
            task[x][y]->post_set[H] = hor;
        }
    }

    // ---- ---- Trigger Execution ---- ---- ----

    // The last task : top-right end
    tbb::task &last_task = *task[begin][end];
    if (VP > 1)
    {
        // When VP > 1, preceding tasks exist. Must wait for them
        last_task.increment_ref_count();
        last_task.spawn_and_wait_for_all(*task[0][0]);
    }
    // Wait done. execute last task
    last_task.execute();

    // ---- ---- Clean-Up/Return ---- ---- ----

    // Explicit memory deallocation
    tbb::task::destroy(last_task);
    return;
}
}
