# Comparison of Go and C++ TBB on Parallel Processing 

This is repository for research codes.   
C++ codes expects Windows developing environment with Visual Studio 2015 or later.

#### Author
[Park Dong-Ha](luncliff@gmail.com) : Dept. of Computer Science & Engineering, Dongguk University 
#### Advisor
Prof. Moon Bong-Kyo : Dept. of Computer Science & Engineering, Dongguk University 

## Abstract
Applying concurrent structure and parallel processing are a common issue for these day’s programs. In this research, Dynamic Programming is used to compare the parallel performance of Go language and Intel C++ Thread Building Blocks. The experiment was performed on 4 core machine and its result contains execution time under Simultaneous Multi-Threading environment. Static Optimal Binary Search Tree was used as an example. 
 
From the result, the speed-up of Go was higher than the number of cores, and that of TBB was close to it. TBB performed better in general, but for larger scale, Go was partially faster than the other.  


## Directory

### 1. `bin/`
For missing DLL files, you can download them from its [official page](https://www.threadingbuildingblocks.org/).

To execute `release` build result, `tbb.dll` is necessary.   
Of course, `debug` build required `tbb_debug.dll`. 

### 2. `lib/`
External libraries
  - [TBB2017](https://github.com/01org/tbb)
  - [CmdParser](https://github.com/FlorianRappl/CmdParser)

### 3. `src/`
Go source files

#### `matrix/`
Minimal Matrix implementation. Referenced [Carl Johnell's research](http://www.diva-portal.org/smash/get/diva2:824741/FULLTEXT03).

#### `obst/`
See [Wikipedia](https://en.wikipedia.org/wiki/Optimal_binary_search_tree) or [this link](http://software.ucv.ro/~mburicea/lab5ASD.pdf)

#### `previous/`
Previous implementation based on ["Comparing parallel performance of Go and C++ TBB on a direct acyclic task graph using a dynamic programming problem"](http://dl.acm.org/citation.cfm?id=2184575&dl=ACM&coll=DL&CFID=748048953&CFTOKEN=40164739).

#### `research/`
Re-implemented code for this research

#### `utils/`
Utility classes

#### `watch/`
Stop watch class over Go Time package

### 4. `build/`
Reserved directory for build result.


## How to Try
The followings are command-line examples. After build, use `-h` for description.

### Go
```
 ./Go.exe -h
 ./Go.exe -n=2048 -np=4 -vp=128 -parallel=true
```

### C++ TBB

### Build
Open `TBB.vcxproj` with Visual Studio 2015 or later, and build with `x64` configuration.
If several TBB files(`.lib`) are missing, download [TBB release](https://github.com/01org/tbb/releases) binaries.

For execution, DLL binary `tbb.dll` is necessary. You can use it in `bin/` folder, but I recommend to download it.
See the [official documentation](https://www.threadingbuildingblocks.org/documentation) first.
```
 ./TBB.exe -h
 ./TBB.exe -n 2048 -np 4 -vp 128 -p true
```

## License
<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.
