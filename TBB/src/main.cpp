// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : main.cpp
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2016/12/17
//
//  Note     :
//      Evaluate Optimal Binary Search Tree problem based on the
//      command-line options.
//      - `N`  : Problem size
//      - `NP` : Number of Processors
//      - `VP` : Scale of Sub-problems
//              Small : big  sub-problem, but low  sync cost
//              Big   : tiny sub-problem, but high sync cost
//  
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#include <iostream>
#include <iomanip>
#include <string>
#include <random>
#include <thread>

#include "./utils/Watch.hpp"    // StopWatch
#include "./Config.hpp"
#include "./Evaluate.hpp"       // Optimal BST evaluation

using namespace std;
using namespace std::chrono;
using namespace Research;

// Alias for stop watch class
using StopWatch = stop_watch<chrono::high_resolution_clock>;

int main(int argc, char* argv[])
{
    Parser parser{ argc, argv };    
    
    // ==== ==== Setup configuration  ==== ====
    
    // Parse the flags
    parser.run_and_exit_if_error(); 
    Config  config  = parser.config();

    // ==== ==== Construct / Initialize ==== ====

    // tree : initialized with random probability value
    auto tree = std::make_unique<Tree>(config.N);
    Init(*tree);

    // ==== ==== ==== Evaluation  ==== ==== ====

    // Start stop watch
    StopWatch       timer{};
    timer.reset();

    // Parallel
    if (config.Parallel == true) {
        // Delimit the number of threads
        tbb::task_scheduler_init sched_init(config.NP);

        EvaluatePar(*tree, config.VP);
    }
    // Sequential 
    else {
        EvaluateSeq(*tree);
    }

    // ==== ==== ==== Result  ==== ==== ====

    auto elapsed = timer.pick<milliseconds>();  // Pick stop watch
    {
        Report report{};
        report.config = config;
        report.elapsed = elapsed.count();

        cout << report << endl; // Print Result
    }

    return EXIT_SUCCESS;
}

