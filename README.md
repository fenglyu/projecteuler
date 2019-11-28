# Start Project Euler


## Run specific unit test
```
go test -timeout 30s github.com/fenglyu/projecteuler/golang/common  -run ^\(TestTailFib\)$
```


## Go pprof 
```
> https://github.com/campoy/go-tooling-workshop/blob/master/3-dynamic-analysis/3-profiling/2-pprof.md
> http://www.graphviz.org/

go tool pprof 27.prof
```


## Go Module proxy
```
export GOPROXY=https://mirrors.aliyun.com/goproxy/
```
