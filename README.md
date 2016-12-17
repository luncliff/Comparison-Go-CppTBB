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
 - [Comparing Parallel Performance of Go and C++ TBB on a Direct Acyclic Task Graph Using a Dynamic Programming Problem](https://pdfs.semanticscholar.org/86dd/6361792684cb844ddeda605ae0e0457efc6c.pdf)

