// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note
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
#include "./research/Config.h"
#include "./research/Evaluate.h"       // Optimal BST evaluation

using namespace std;
using namespace std::chrono;
using namespace Research;

int main(int argc, char* argv[])
{
    Parser parser{ argc, argv };    
    
    // ==== ==== Setup configuration  ==== ====
    
    // Parse the flags
    parser.run_and_exit_if_error(); 
    Config  config  = parser.config();

    // ==== ==== Construct / Initialize ==== ====

    // Optimal Binary Search Tree
    //      Allocated on heap for memory profiling
    auto tree = std::make_unique<Tree>(config.N);
    // Assign random probability value
    Init(*tree);  

    // ==== ==== ==== Evaluation  ==== ==== ====

    StopWatch       timer{};
    timer.reset();

    // Parallel
    if (config.Parallel == true) {
        // Delimit the number of threads
        tbb::task_scheduler_init scheduler(config.NP);

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

