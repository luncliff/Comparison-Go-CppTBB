// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note     :
//      Configuration + Command argument parser
//  
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _RESEARCH_CONFIG_HPP_
#define _RESEARCH_CONFIG_HPP_

#include <iomanip>
#include <thread>
#include "./Alias.h"
#include <cmdparser.hpp>  // https://github.com/FlorianRappl/CmdParser

namespace Research
{
    // - Note
    //      Experiment Configuration
    //      - `N`  : Problem size
    //      - `NP` : Number of Processors
    //      - `VP` : Scale of Sub-problems
    //              Small : big  sub-problem, but low  sync cost
    //              Big   : tiny sub-problem, but high sync cost
    struct Config
    {
        i32 N, NP, VP;
        bool Parallel;
    };

    // - Note
    //      Print to stream with JSON format
    std::ostream& operator <<(std::ostream& out, const Config& cfg);


    // ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----


    class Parser : 
            public cli::Parser
    {
    public:

        // - Note
        //      Parser setup
        Parser(int argc, char* argv[]);


        Config config() const noexcept(false);
    };



    struct Report
    {
        Config  config;     // Configuration
        i64     elapsed;    // Elapsed time
    };

    // - Note
    //      Print to stream with JSON format
    std::ostream& operator <<(std::ostream& out, const Report& rep);



}
#endif
