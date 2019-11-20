# Start Project Euler


## run unit test
```
go test -timeout 30s github.com/fenglyu/projecteuler/golang/common  -run ^\(TestTailFib\)$
```

## pprof 
> https://github.com/campoy/go-tooling-workshop/blob/master/3-dynamic-analysis/3-profiling/2-pprof.md
> http://www.graphviz.org/
```
go tool pprof 27.prof
```

## Go Module proxy
```
1.使用go1.11以上版本并开启go module机制

2.导出GOPROXY环境变量

export GOPROXY=https://mirrors.aliyun.com/goproxy/
```
