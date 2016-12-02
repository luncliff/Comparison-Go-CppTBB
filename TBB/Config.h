#pragma once

struct Config
{
    int N, NP, VP;

    Config() = default;
    Config(int n, int np, int vp) :
        N{ n }, NP{ np }, VP{ vp }
    {}
};
