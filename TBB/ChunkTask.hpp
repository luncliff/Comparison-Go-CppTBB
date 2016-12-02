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
	// No preset. Post set only
	tbb::task* post_set[2];
    
public:
    // ctor
	ChunkTask(Tree& _tree, 
              int _i, int _j, int _vp) :
        tree{ _tree },
        i {_i}, j{ _j },   vp{ _vp }
	{}

    // Processing chunk
	tbb::task* execute() override 
	{
		this->wait(); 

        // ---- ---- Processing ---- ----
        // loop: bottom-left >>> top-right
        for (auto row = vp - 1; row >= 0; --row) {
            for (auto col = 0; col < vp - 1; ++col) {
                // Calculate without side-effect
                auto root_cost = Tree::Calculate(tree, row, col);
                // Assign results
                tree.root[row][col] = std::get<0>(root_cost);
                tree.cost[row][col] = std::get<1>(root_cost);
            }
        }

        // Calculation
        //for (auto ii = ih; ii >= il; --ii) {
        //    if (i == j) {
        //        bb == ii;
        //    }
        //    else {
        //        bb = jl;
        //    }
        //    for (auto jj = bb; jj <= jh; ++jj) {
        //        auto tuple = obst.calculate(ii, jj);
        //        obst.root[ii][jj] = std::get<0>(tuple);
        //        obst.cost[ii][jj] = std::get<1>(tuple);
        //    }
        //}


		// ---- ---- ---- ---- ---- ----
		
        this->notify();	
		return nullptr;	// No task chaining
	}

    // Wait for pre-set
	void wait() 
    {
		tbb::task::wait_for_all();
	}
	
    // Notify to post-set
	void notify() 
	{
		if (post_set[0] != nullptr) {   
            post_set[0]->decrement_ref_count(); 
        }
		if (post_set[1] != nullptr) {   
            post_set[1]->decrement_ref_count(); 
        }
	}
};
