// ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ==== ====
//
//  Author   : Park  Dong Ha ( luncliff@gmail.com )
//
//  Note     :
//      Evaluate Optimal Binary Search Tree problem.
//      - `EvaluateSeq` : Sequential processing
//      - `EvaluatePar` : Parallel processing with Intel TBB
//
// ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ---- ----
#ifndef _RESEARCH_EVALUATE_HPP_
#define _RESEARCH_EVALUATE_HPP_

#include "./Alias.h"
#include "./ChunkTask.hpp"
#include "./Config.h"
#include "../utils/Matrix.h"

namespace Research
{

// - Note
//      Sequential Evaluation
void EvaluateSeq(Tree &tree) noexcept;

// - Note
//      Parallel Evaluation
//      Chained spawning, Explicit task destroy
// - Reference
//      https://software.intel.com/en-us/node/506110
void EvaluatePar(Tree &tree, const i32 VP) noexcept;
}
#endif
