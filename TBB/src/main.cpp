// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : main.cpp
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2016/12/05
//
//  Note     :
//      So tired....
//  
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#include <iostream>
#include <iomanip>
#include <string>
#include <random>
#include <thread>

#include <CmdParser.hpp>

#include "./utils/Watch.hpp"
#include "./Evaluate.hpp"

using StopWatch = magic::stop_watch<std::chrono::high_resolution_clock>;

static void Setup(cli::Parser& _parser);

int main(int argc, char* argv[])
{
    using namespace std;
    using namespace std::chrono;
    using namespace Research;

    cli::Parser parser{argc, argv};
    Setup(parser);
    parser.run_and_exit_if_error(); // Parse the flags
    
    // ---- ---- Setup config ---- ---- 
    bool is_par = (parser.get<std::string>("p") == "true") ? 
                      true : false;

    Config  cfg{};
    cfg.N = parser.get<int>("n");
    cfg.NP = parser.get<int>("np");
    cfg.VP = parser.get<int>("vp");

    if (is_par == false) {
        cfg.NP = 1;
    }
    Display(cfg);


    // ---- Construct / Initialize ----
    auto tree = std::make_unique<Tree>(cfg.N);
    Init(*tree);

    // ---- ---- Evaluation  ---- ----
    StopWatch       watch{};
    if (is_par == true) {
        // Parallel
        EvaluatePar(cfg, *tree);
    }
    else {
        // Sequential 
        EvaluateSeq(cfg, *tree);
    }

    // ---- ---- Result ---- ----
    auto dur = watch.pick<milliseconds>();
    printf_s(" [ %10s ] : %8d ms \n",
            (is_par)? "Parallel" : "Sequential",
            dur.count());

    return EXIT_SUCCESS;
}



static void Setup(cli::Parser& _p) 
{
    using namespace std;

    // Default values...
    int N = 1 << 11;
    int NP = thread::hardware_concurrency();
    int VP = NP*NP;

    // Set options...
    _p.set_optional<int>("n",  "size",     N,
                         "Problem's size");
    _p.set_optional<int>("np", "proc",     NP,
                         "Number of physical processor");
    _p.set_optional<int>("vp", "chunk",    VP,
                         "Sub-problem's size");
    _p.set_optional<std::string>("p", "parallel", "true",
                          "Parallel execution");
    return;
}



