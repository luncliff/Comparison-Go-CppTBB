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

## License
<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.


## How to Try
The followings are command-line examples. After build, use `-h` for description.

### Go
```
 ./Go.exe -h
 ./Go.exe -n=2048 -np=4 -vp=128 -parallel=true
```

### C++ TBB
For `release` build's execution, DLL binary `tbb.dll` is necessary. 
See the [official documentation](https://www.threadingbuildingblocks.org/documentation) first.
```
 ./TBB.exe -h
 ./TBB.exe -n 2048 -np 4 -vp 128 -p true
```



