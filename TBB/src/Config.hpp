// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  File     : Config.h
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//  Updated  : 2016/12/17
//
//  Note     :
//      Configuration + Command argument parser
//  
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _RESEARCH_CONFIG_HPP_
#define _RESEARCH_CONFIG_HPP_

#include <iomanip>

#include "./Alias.h"
#include <CmdParser.hpp>  // https://github.com/FlorianRappl/CmdParser

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
    static 
    std::ostream& operator <<(std::ostream& out, const Config& cfg)
    {
        using namespace std;
        constexpr char Format[] =
            "{ \"N\" : %d, \"Proc\" : %d, \"VP\" : %d, \"Parallel\" : %s }";

        char buf[64]{};
        sprintf_s(buf, 
            Format, cfg.N, cfg.NP, cfg.VP, cfg.Parallel? "true": "false");

        return out << buf;
    }


    // ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----


    class Parser : 
            public cli::Parser
    {
    public:

        // - Note
        //      Parser setup
        Parser(int argc, char* argv[]) : cli::Parser{ argc,argv }
        {            
            int N = 1 << 11;    // Fixed size : 2048
            int NP = std::thread::hardware_concurrency();
            int VP = NP*NP;     // Square of NP

            // Set options...
            set_optional<int>("n", "size", N,
                        "Problem's size");
            set_optional<int>("np", "proc", NP,
                        "Number of physical processor");
            set_optional<int>("vp", "chunk", VP,
                        "Sub-problem's size");
            set_optional<std::string>("p", "parallel", "true",
                        "Parallel execution");
        }


        Config config() const noexcept(false)
        {
            Config cfg{};

            cfg.N           = this->get<int>("n");
            cfg.NP          = this->get<int>("np");
            cfg.VP          = this->get<int>("vp");
            cfg.Parallel    = get<std::string>("p") == "true";

            // If sequential...
            if (cfg.Parallel == false) {
                cfg.NP = 1;
                cfg.VP = 1;
            }

            return cfg;
        }
    };



    struct Report
    {
        Config  config;     // Configuration
        i64     elapsed;    // Elapsed time
    };

    // - Note
    //      Print to stream with JSON format
    static
    std::ostream& operator <<(std::ostream& out, const Report& rep) 
    {
        return out
            << "{ "
            << "\"Config\" : " << rep.config << ", "
            << "\"Elapsed\" : " << rep.elapsed
            << " }";
    }



}
#endif
