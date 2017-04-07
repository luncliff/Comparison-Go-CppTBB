// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Note
//      Previous research's implementation
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _RESEARCH_CHUNK_TASK_HPP_
#define _RESEARCH_CHUNK_TASK_HPP_

// - Dependency
//      Intrl C++ Thread Building Blocks
#include <tbb/task.h>
#include <tbb/task_scheduler_init.h>

#include "../research/Tree.h"

static void chunk(Tree &tree, int i, int j, int vp, int n)
{
    auto il = (i * (n + 1)) / vp;
    auto jl = (j * (n + 1)) / vp;
    auto ih = (((i + 1) * (n + 1)) / vp) - 1;
    auto jh = (((j + 1) * (n + 1)) / vp) - 1;

    for (i32 row = ih; row >= il; --row)
    {
        i32 col = (i == j) ? row : jl;
        for (; col <= jh; ++col)
        {
            // Calculate without side-effect
            auto root_cost = Tree::Calculate(tree, row, col);
            // Assign results
            tree.root[row][col] = std::get<0>(root_cost);
            tree.cost[row][col] = std::get<1>(root_cost);
        }
    }
}

class DagTask : 
    public tbb::task
{
    const int i, j, vp, n;
    Tree &tree;
  public:
    DagTask *successor[2]{};
    DagTask *a{};
    //next task of the front
    DagTask(Tree &tree, 
            int i_, int j_, int vp_, int n_) noexcept :
        tree{tree},
        i(i_), j(j_), vp(vp_), n(n_)
    {}

    tbb::task *execute()
    {
        if (i == j && i != vp - 1)
        {
            //main diagonal
            spawn(*a); //front node
        }

        chunk(tree, i, j, vp, n);

        for (int k = 0; k < 2; ++k)
        {
            if (DagTask *t = successor[k])
                if (t->decrement_ref_count() == 0)
                    spawn(*t);
        }
        return nullptr;
    }
};

#endif
