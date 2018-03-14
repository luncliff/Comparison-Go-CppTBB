#include <sdkddkver.h>
#include "CppUnitTest.h"

using namespace Microsoft::VisualStudio::CppUnitTestFramework;

#include "../src/obst.h"

namespace CppTBB
{		
	TEST_CLASS(Sequential)
	{
	public:

		TEST_METHOD(Test1)
		{
            const uint32_t Width = 32;

            obst tree{ Width };

            const uint32_t N = static_cast<uint32_t>(tree.size());
            Assert::IsTrue(Width == N);
 
            tree.randomize();

            // loop : bottom-left >>> top-right
            //      [ + + + + ]
            //      [   + + + ]
            //      [     + + ]
            //   -> [       + ]
            for (int32_t r = N; r >= 0; --r)
            {
                for (uint32_t c = r; c <= N; ++c)
                {
                    // Calculate without side-effect
                    const auto optimal = obst::optimal(tree, r, c);

                    // Assign result
                    tree.root[r][c] = std::get<0>(optimal);
                    tree.cost[r][c] = std::get<1>(optimal);
                }
            }

            Assert::IsTrue(Width == N);
        }

	};
}