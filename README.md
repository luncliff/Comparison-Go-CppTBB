# Comparison of Go and C++ TBB on Parallel Processing 

## Summary
This is the personal research about Golang and C++ TBB.   
Visit the [Wiki](https://github.com/luncliff/Research-Go-Cpp/wiki) for detail.

## How to test
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

## License
<a rel="license" href="http://creativecommons.org/licenses/by/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by/4.0/">Creative Commons Attribution 4.0 International License</a>.


