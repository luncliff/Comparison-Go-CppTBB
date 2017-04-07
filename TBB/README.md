## Directory
### `bin/`
For missing DLL files, you can download them from its [official page](https://www.threadingbuildingblocks.org/).

To execute `release` build result, `tbb.dll` is necessary.   
Of course, `debug` build required `tbb_debug.dll`. 

### `lib/`
External libraries
  - [TBB2017](https://github.com/01org/tbb)
  - [CmdParser](https://github.com/FlorianRappl/CmdParser)

### `build/`
Reserved directory for build result.

### `src/`
#### `previous/`
Previous implementation based on ["Comparing parallel performance of Go and C++ TBB on a direct acyclic task graph using a dynamic programming problem"](http://dl.acm.org/citation.cfm?id=2184575&dl=ACM&coll=DL&CFID=748048953&CFTOKEN=40164739).

#### `research/`
Re-implemented code for this research

#### `utils/`
Utility classes
