# Comparison of Go and C++ TBB on Parallel Processing 

## Summary
This is the personal research about Golang and C++ TBB.   
Visit the [Wiki](https://github.com/luncliff/Research-Go-Cpp/wiki) for detail.

## How to test
The followings are command-line examples

### Go
```
 ./Go.exe -n=2048 -np=4 -vp=128
```

### TBB
For `release` build's execution, `tbb.dll` is necessary.
```
 ./TBB.exe -n "2048" -np "4" -vp "128" 
```

## Reference
 - Doug Serfass, Peiyi Tang. [Comparing parallel performance of Go and C++ TBB on a direct acyclic task graph using a dynamic programming problem](http://dl.acm.org/citation.cfm?id=2184575), 2012 March, ACM
