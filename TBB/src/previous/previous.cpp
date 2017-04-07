// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : previous.cpp
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2017/02/03
//
//  Note     :
//      Evaluate Optimal Binary Search Tree problem
//  
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#include <iostream>
#include <iomanip>
#include <string>
#include <random>
#include <thread>

#include "../utils/Watch.hpp"    // StopWatch
#include "../research/Config.h"
#include "./DAGTask.hpp"       // Optimal BST evaluation

using namespace std;
using namespace std::chrono;
using namespace Research;

// Alias for stop watch class
using StopWatch = stop_watch<chrono::high_resolution_clock>;

void EvaluatePar(Tree& tree, u32 vp);
void EvaluateSeq(Tree& tree);


namespace test
{
    // ==== ==== Setup configuration  ==== ====

    constexpr u32 N = 1 << 11;
    constexpr u32 NP = 4;
    constexpr u32 VP = NP * NP;
    constexpr bool Parallel = true;
}

Tree tree{ static_cast<size_t>(test::N) };

int previous_main(int argc, char* argv[])
{
    // ==== ==== Construct / Initialize ==== ====

    // tree : initialized with random probability value
    Init(tree);

    // ==== ==== ==== Evaluation  ==== ==== ====

    // Start stop watch
    StopWatch       timer{};
    timer.reset();

    // Parallel
    if (test::Parallel == true) {
        // Delimit the number of threads
        tbb::task_scheduler_init scheduler(test::NP);

        EvaluatePar(tree, test::VP);
    }
    // Sequential 
    else {
        EvaluateSeq(tree);
    }

    // ==== ==== ==== Result  ==== ==== ====
    //Display(tree);

    auto elapsed = timer.pick<milliseconds>();  // Pick stop watch
    {
        Report report{};
        report.config.N = test::N;
        report.config.NP = test::NP;
        report.config.VP = test::VP;
        report.config.Parallel = test::Parallel;
        report.elapsed = elapsed.count();

        cout << report << endl; // Print Result
    }

    return EXIT_SUCCESS;
}


void EvaluateSeq(Tree& tree) {
    const i32 n = static_cast<i32>(tree.size());
    for (auto i = n; i >= 0; --i) {
        for (auto j = i; j <= n; j++) {

            std::tie(tree.root[i][j], tree.cost[i][j])
                = tree.Calculate(tree, i, j);
        }
    }
}

// - Note
//      Previous research's implementation
void EvaluatePar(Tree& tree, u32 vp)
{
    const int n = static_cast<int>(tree.size());
    Matrix<DagTask*> x{ vp,vp };

    //allocate,initialize arrays cost,root,prob . 
    //create tasks 
    for (u32 d = 0; d < vp; d++)
    {
        for (u32 i = 0; i + d < vp; i++) {
            auto proxy = tbb::task::allocate_root();
            x[i][i + d] = new(proxy)
                DagTask(tree, i, i + d, vp, n);

            x[i][i + d]->set_ref_count(d == 0 ? 0 : 2);
        } 
        //set up successor links and front link 
    }

    for (u32 d = 0; d < vp; d++)
    {
        for (int i = 0; i + d < vp; i++) {
            x[i][i + d]->successor[0] = i - 1 > -1 ? x[i - 1][i + d] : nullptr;// up 
            x[i][i + d]->successor[1] = i + d + 1<vp ? x[i][i + d + 1] : nullptr;// right 

            if (d == 0 && static_cast<u32>(i) < vp - 1) //main diagonal 
                x[i][i + d]->a = x[i + 1][i + 1];
        }
    }

    if (vp > 1) {
        x[0][vp - 1]->increment_ref_count();
        x[0][vp - 1]->spawn_and_wait_for_all(*x[0][0]);
    }
    x[0][vp - 1]->execute();

    tbb::task::destroy(*x[0][vp - 1]);
    return;
}

