#include "./Config.h"

namespace Research
{
// - Note
//      Print to stream with JSON format
std::ostream &operator<<(std::ostream &out, const Config &cfg)
{
    using namespace std;
    constexpr char Format[] =
        "{ \"N\" : %d, \"Proc\" : %d, \"VP\" : %d, \"Parallel\" : %s }";

    char buf[64]{};
    sprintf_s(buf,
              Format, cfg.N, cfg.NP, cfg.VP, cfg.Parallel ? "true" : "false");

    return out << buf;
}

Parser::Parser(int argc, char *argv[]) : cli::Parser{argc, argv}
{
    int N = 1 << 11; // Fixed size : 2048
    int NP = std::thread::hardware_concurrency();
    int VP = NP * NP; // Square of NP

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

Config Parser::config() const noexcept(false)
{
    Config cfg{};

    cfg.N = this->get<int>("n");
    cfg.NP = this->get<int>("np");
    cfg.VP = this->get<int>("vp");
    cfg.Parallel = get<std::string>("p") == "true";

    // If sequential...
    if (cfg.Parallel == false)
    {
        cfg.NP = 1;
        cfg.VP = 1;
    }

    return cfg;
}

std::ostream &operator<<(std::ostream &out, const Report &rep)
{
    return out
           << "{ "
           << "\"Config\" : " << rep.config << ", "
           << "\"Elapsed\" : " << rep.elapsed
           << " }";
}
}