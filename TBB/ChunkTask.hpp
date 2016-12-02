#ifndef _CHUNK_TASK_HPP_
#define _CHUNK_TASK_HPP_

#include <tbb/task.h>
#include "Tree.h"


class ChunkTask : public tbb::task
{
private:
    Tree&   tree;
    int     i, j;   // top-left index of chunk
    int     vp;     // chunk width

public:
    // Chained task
    tbb::task* chain = nullptr;
	// No preset. Post set only
    tbb::task* post_set[2]{};
    
public:
    // ctor
	ChunkTask(Tree& _tree, 
              int _i, int _j, int _vp) :
        tree{ _tree },
        i {_i}, j{ _j },   vp{ _vp }
	{}

    // Wait for pre-set
    //      In this implementation, Intel TBB scheduler will handle this.
    //void wait();

    // Notify to post-set
    //      Decrement the successors' ref count.
    //      If possible, enqueue them into scheduler
    void notify()
    {
        if (post_set[0] != nullptr) {
            // if ready, count will be 0
            if (post_set[0]->decrement_ref_count() == 0) {
                // spawn : enqueue to scheduler
                spawn(*post_set[0]);
            }
        }
        if (post_set[1] != nullptr) {
            // if ready, count will be 0
            if (post_set[1]->decrement_ref_count() == 0) {
                // spawn : enqueue to scheduler
                spawn(*post_set[1]);
            }
        }
    }


    // Processing chunk
	tbb::task* execute() override 
	{
		// No waiting.
        // if there is chained task, spawn it immediately
        if (chain != nullptr) { spawn(*chain); }


        // ---- ---- Processing ---- ----
        // loop: bottom-left >>> top-right
        for (auto row = vp - 1; row >= 0; --row) {
            for (auto col = 0; col < vp - 1; ++col) 
            {
                // Calculate without side-effect
                auto root_cost = Tree::Calculate(tree, row, col);
                // Assign results
                tree.root[row][col] = std::get<0>(root_cost);
                tree.cost[row][col] = std::get<1>(root_cost);
            }
        }
		// ---- ---- ---- ---- ---- ----


        // Notify successors and enqueue to scheduler
        this->notify();	
        // No task chaining
		return nullptr;	
	}


};
#endif
