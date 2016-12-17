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

#include <CmdParser.hpp>        // Command Line Parser

#include "./utils/Watch.hpp"    // StopWatch
#include "./Evaluate.hpp"       // Optimal BST evaluation

using namespace std;
using namespace std::chrono;
using namespace Research;

// Alias for stop watch class
using StopWatch = stop_watch<chrono::high_resolution_clock>;

// Parser setup
static cli::Parser make_parser(int argc, char* argv[]);


int main(int argc, char* argv[])
{
    cli::Parser p = make_parser(argc, argv);
    p.run_and_exit_if_error(); // Parse the flags
    
    // ---- ---- Setup configuration ---- ---- 
    bool parallel = (p.get<std::string>("p") == "true") ? 
                      true : false;

    Config  cfg{};
    {
        cfg.N  = p.get<int>("n");
        cfg.NP = p.get<int>("np");
        cfg.VP = p.get<int>("vp");

        if (parallel == false) {
            cfg.NP = 1;
            cfg.VP = 0;
        }
        Display(cfg);
    }

    // ---- Construct / Initialize ----
    auto tree = std::make_unique<Tree>(cfg.N);
    Init(*tree);

    // ---- ---- Evaluation  ---- ----
    StopWatch       watch{};    // Start stop watch
    if (parallel == true) {
        // Parallel
        EvaluatePar(cfg, *tree);
    }
    else {
        // Sequential 
        EvaluateSeq(cfg, *tree);
    }

    // ---- ---- Result ---- ----
    auto dur = watch.pick<milliseconds>();  // Pick stop watch
    printf_s(" [ %10s ] : %8d ms \n",
            (parallel)? "Parallel" : "Sequential",
            dur.count());

    return EXIT_SUCCESS;
}



static cli::Parser make_parser(int argc, char* argv[])
{
    using namespace std;

    cli::Parser parser{ argc, argv };

    // Fixed size : 2048
    int N = 1 << 11;    
    // follow standard
    int NP = thread::hardware_concurrency();    
    // Square of NP
    int VP = NP*NP; 

    // Set options...
    parser.set_optional<int>("n",  "size",     N,
                         "Problem's size");
    parser.set_optional<int>("np", "proc",     NP,
                         "Number of physical processor");
    parser.set_optional<int>("vp", "chunk",    VP,
                         "Sub-problem's size");
    parser.set_optional<string>("p", "parallel", "true",
                          "Parallel execution");
    return parser;
}



