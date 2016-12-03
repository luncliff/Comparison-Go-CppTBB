
#include <iostream>
#include <iomanip>
#include <random>

#include "./utils/Watch.hpp"
#include "./Evaluate.hpp"

using StopWatch = magic::stop_watch<std::chrono::high_resolution_clock>;

static void Display(const Config& _cfg);
static void Display(const Tree& _tree);

static void Setup(const Config& _cfg) 
{
    tbb::task_scheduler_init sched_init(_cfg.NP);
}

static void Init(Tree& _tree);

int main(void)
{
    using namespace std;
    using namespace std::chrono;

    Config  config{};
    config.N  = 1 << 11;
    config.NP = 1 << 3;
    config.VP = 1 << 3;

    Setup(config);
    Display(config);

    // ---- Construct / Initialize ----
    auto tree = std::make_unique<Tree>(config.N);
    Init(*tree);

    cout << "Initialization done." << endl;
    // ---- ---- Evaluation  ---- ----
    StopWatch watch{};

    EvaluatePar(config, *tree);

    auto dur = watch.reset<milliseconds>();
    // ---- ---- ---- ---- ---- ----
    
    //Display(*tree);
    cout << dur.count() << " ms" << endl;

    return std::system("pause");
}



static void Init(Tree& _tree)
{
    using namespace std;
    using namespace std::chrono;

    auto duration = system_clock::now().time_since_epoch();
    auto seed = static_cast<u32>(duration.count());
    mt19937 rnd{ seed };

    f64     sum{};
    for (i32 i = 0; i < _tree.prob.size(); ++i) {
        sum += _tree.prob[i] = rnd() % (1 << 20);
    }
    for (i32 i = 0; i < _tree.prob.size(); ++i) {
        _tree.prob[i] /= sum;
    }
}



static void Display(const Config& _cfg)
{
    using namespace std;

    printf_s(" [ Proc ] : %5d\n", _cfg.NP);
    printf_s(" [ N    ] : %5d\n", _cfg.N);
    printf_s(" [ VP   ] : %5d\n", _cfg.VP);
}

static void Display(const Tree& _tree)
{
    using namespace std;

    const i32 N = static_cast<i32>(_tree.size());

    puts("--- Prob ---");
    for (i32 i = 0; i < N; ++i) {
        printf_s(" %4.4lf, ", _tree.prob[i]);
    }
    putchar('\n');

    puts("--- Root/Cost ---");
    for (i32 i = 0; i < N + 1; ++i) {
        for (i32 j = 0; j < N + 1; ++j)
        {
            auto root = _tree.root[i][j];
            auto cost = _tree.cost[i][j];
            printf_s(" [%2d,%2.2lf] ", root, cost);
        }
        putchar('\n');
    }
}
