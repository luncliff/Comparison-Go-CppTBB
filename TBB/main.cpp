#include <iostream>

// #include watch
#include "Evaluate.hpp"

using namespace std;
using namespace tbb;

int main() 
{
    Config  cfg{};
    Tree    tree{ cfg.N };

    EvaliateSeq(cfg, tree);

    return std::system("pause");
}


