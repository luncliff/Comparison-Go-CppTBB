#pragma once
#include <cstdint>

using u32 = std::uint32_t;
using i32 = std::int32_t;
using u64 = std::uint64_t;

struct Config
{
    u32 N, NP, VP;

    Config() = default;
    Config(u32 n, u32 np, u32 vp) :
        N{ n }, NP{ np }, VP{ vp }
    {}
};
