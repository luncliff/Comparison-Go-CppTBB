#include "obst.h"

obst::obst(std::size_t N) noexcept(false) :
    prob(N),
    cost{ N + 1 },
    root{ N + 1 }
{
}

size_t obst::size() const noexcept
{
    return prob.size();
}

void obst::randomize() noexcept
{
    using namespace std;
    using namespace std::chrono;

    // Use epoch time as the seed of random generator
    auto duration = system_clock::now().time_since_epoch();
    auto seed = static_cast<uint32_t>(duration.count());
    mt19937 rnd{ seed };

    float total = 0;
    // set random value and accomulate total
    for (float& p : prob) {
        p = static_cast<float>(rnd() % (1 << 20));
        total += p;
    }

    // normalize
    for (float& p : prob) 
        p /= total;

    return;
}

auto obst::optimal(const obst& tree,
    uint32_t _row, uint32_t _col) noexcept
    ->std::tuple<int32_t, float>
{
    int32_t best_root = -1;
    float best_weight = FLT_MAX;
    float weight = 0;

    // Unused range
    if (_row >= _col)
    {
        weight = 0;
    }
    // Main diagonal
    else if (_row + 1 == _col)
    {
        best_root = _row + 1;
        weight = tree.prob[_row];
    }
    // Tree estimation
    else
    {
        float sum = 0; // basic weight
        for (auto i = _row; i < _col; ++i)
        {
            // Accumulate
            sum += tree.prob[i];

            // Find best weight
            float temp_weight =
                tree.cost[_row][i] + tree.cost[i + 1][_col];

            if (temp_weight < best_weight)
            {
                best_weight = temp_weight;
                best_root = i + 1;
            }
        }
        weight = best_weight + sum;
    }
    // Weight == Cost
    return std::make_tuple(best_root, weight);
}


bool operator==(const obst &lhs, const obst &rhs) noexcept
{
    if (lhs.prob != rhs.prob)
    {
        return false;
    }
    const auto N = lhs.size();

    // Compare valid range(x) only.
    //   [ - x x x ]
    //   [ - - x x ]
    //   [ - - - x ]
    //   [ - - - - ]
    for (auto i = 0u; i < N; ++i)
    {
        for (auto j = i + 1; j < N; ++j)
        {
            const bool equal_root = lhs.root[i][j] != rhs.root[i][j];
            const bool equal_cost = lhs.cost[i][j] != rhs.cost[i][j];

            if (equal_root == false || equal_cost == false)
                return false;

        }
    }
    return true;
}


bool operator!=(const obst& lhs, const obst& rhs) noexcept
{
    return !(lhs == rhs);
}




//
//// - Note
////      Display the tree in console.
////      Used for debugging
//void Display(const Tree &_tree)
//{
//    using namespace std;
//
//    const i32 N = static_cast<i32>(_tree.size());
//
//    puts("--- Prob ---");
//    for (i32 i = 0; i < N; ++i)
//    {
//        printf_s(" %4.4lf, ", _tree.prob[i]);
//    }
//    putchar('\n');
//
//    puts("--- Root/Cost ---");
//    for (i32 i = 0; i < N + 1; ++i)
//    {
//        for (i32 j = 0; j < N + 1; ++j)
//        {
//            auto root = _tree.root[i][j];
//            auto cost = _tree.cost[i][j];
//            printf_s(" [%2lld,%2.2lf] ", root, cost);
//        }
//        putchar('\n');
//    }
//}
