// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : ChunkTask.hpp
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2016/12/17
//
//  Note     :
//      The sub-problems are grouped to a chunk, and each goroutine
//		will process given single chunk.
//      Process chunks with task-based parallism. 
//
//  Concept  : 
//      - Main Problem : 
//          Evaluating a `Tree`. 
//          For parallel processing, it is divided into sub-problems.
//      - Sub-problem : 
//          Calculating `root` and `cost` with given vertices.
//      - Chunk
//          To process efficiently, sub-problems are *chunked*.
//          The size of chunk can be small so each chunks are mapped to 
//          sub-problem in 1:1 relation. (VP==N)
//          Or, it can be big to reduce synchronization overhead. 
//          (VP << N)
//  
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _RESEARCH_CHUNK_TASK_HPP_
#define _RESEARCH_CHUNK_TASK_HPP_

// - Dependency
//      Intrl C++ Thread Building Blocks
#include <tbb/task.h>
#include <tbb/task_scheduler_init.h>

#include "./Tree.h"

// - Note
//      Task to process chunk.
class ChunkTask : public tbb::task
{
private:
    Tree&   tree;   // Reference for OBST
    i32     i, j;   // Top-left index of chunk
    i32     width;  // Chunk's width

public:
    tbb::task* chain = nullptr;     // Chained task
    tbb::task* post_set[2]{};       	// No preset. Post set only
    
public:
    // - Note
    //      Remember tree and chunk's information
	ChunkTask(Tree& _tree, 
              i32 _i, i32 _j, i32 _width) :
        tree{ _tree },
        i {_i}, j{ _j }, width{ _width }
	{}

    // - Note
    //      Wait for pre-set.
    //      In this implementation, Intel TBB scheduler 
    //      will handle this issue automatically
    //void wait();

    // - Note
    //      Notify to post-set.
    //      Decrement the successors' ref count.
    //      If possible, enqueue them into scheduler
    void notify()
    {
        // if ready, ref count will become 0
        if (post_set[0] != nullptr 
            && post_set[0]->decrement_ref_count() == 0) 
        {
            // spawn : enqueue to TBB scheduler
            spawn(*post_set[0]);
        }

        // if ready, ref count will become 0
        if (post_set[1] != nullptr 
            && post_set[1]->decrement_ref_count() == 0) 
        {
            // spawn : enqueue to TBB scheduler
            spawn(*post_set[1]);
        }
    }

    // - Note
    //      Process the given chunk sequentially.
    //      When ready, TBB scheduler will invoke this function 
    //      After processing, notify to successor tasks...
	tbb::task* execute() override 
	{
        // wait();		// No waiting.

        // if there is chained task, spawn it immediately
        if (chain != nullptr) { spawn(*chain); }

        // ---- ---- Processing ---- ----
        // In range of chunk, process sequentially.
        // loop: bottom-left >>> top-right
        //       [ . . . . ]
        //       [ 9 . . . ]
        //       [ 5 6 . . ]
        //    -> [ 1 2 3 4 ]
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

        // No task bypass
		return nullptr;	
	}

};
#endif
