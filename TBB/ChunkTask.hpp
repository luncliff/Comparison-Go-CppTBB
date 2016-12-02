#ifndef _RESEARCH_CHUNK_TASK_HPP_
#define _RESEARCH_CHUNK_TASK_HPP_

#include <tbb/task.h>
#include <tbb/task_scheduler_init.h>

#include "./Tree.h"

class ChunkTask : public tbb::task
{
private:
    Tree&   tree;
    i32     i, j;   // top-left index of chunk
    i32     width;     // chunk width

public:
    // Chained task
    tbb::task* chain = nullptr;
	// No preset. Post set only
    tbb::task* post_set[2]{};
    
public:
    // ctor
	ChunkTask(Tree& _tree, 
              i32 _i, i32 _j, i32 _width) :
        tree{ _tree },
        i {_i}, j{ _j }, width{ _width }
	{}

    // Wait for pre-set
    //      In this implementation, Intel TBB scheduler will handle this.
    //void wait();

    // Notify to post-set
    //      Decrement the successors' ref count.
    //      If possible, enqueue them into scheduler
    void notify()
    {
        // if ready, ref count will become 0
        if (post_set[0] != nullptr 
            && post_set[0]->decrement_ref_count() == 0) 
        {
            // spawn : enqueue to scheduler
            spawn(*post_set[0]);
        }

        // if ready, ref count will become 0
        if (post_set[1] != nullptr 
            && post_set[1]->decrement_ref_count() == 0) 
        {
            // spawn : enqueue to scheduler
            spawn(*post_set[1]);
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
        for (i32 row = i-1 + width; row >= i; --row) {
            for (i32 col = j; col < j + width; ++col)
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
