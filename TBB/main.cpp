#include <iostream>
#include <iomanip>
#include <tbb/task_scheduler_init.h>
// #include watch
#include "Evaluate.hpp"

using namespace std;
using namespace tbb;

std::ostream& operator << (std::ostream& _out, const Tree& _tree);

void init(const Config& _cfg) 
{
    //tbb::task_scheduler_init sched_init{};
    //sched_init.initialize(_cfg.NP);

}

int main(void)
{
    Config  cfg{};
    cfg.N  = 1 << 2;
    cfg.NP = 1 << 0;
    cfg.VP = 1 << 1;

    init(cfg);

    auto tree = std::make_unique<Tree>(cfg.N);
    
    tree->prob[0] = 0.1*1;
    tree->prob[1] = 0.1*2;
    tree->prob[2] = 0.1*3;
    tree->prob[3] = 0.1*4;

    EvaluatePar(cfg, *tree);

    std::cout << *tree << std::endl;

    return std::system("pause");
}


std::ostream& operator << (std::ostream& _out, const Tree& _tree)
{
    using namespace std;

    const i32 N = _tree.size();

    puts("--- Prob ---");
    for (i32 i = 0; i < N; ++i) {
        _out << _tree.prob[i] << ", ";
    }
    _out << '\n';

    puts("--- Root/Cost ---");
    for (i32 i = 0; i < N + 1; ++i) {
        for (i32 j = 0; j < N + 1; ++j)
        {
            auto root = _tree.root[i][j];
            auto cost = _tree.cost[i][j];
            printf(" [%2d,%2.2lf] ", root, cost);
        }
        _out << "\n";
    }
    return _out;
}
