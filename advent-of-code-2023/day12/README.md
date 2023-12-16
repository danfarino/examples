```
$ ./run.sh
goos: linux
goarch: amd64
pkg: github.com/danfarino/examples/advent-of-code-2023/day12
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkSlow-8           	       1	1239590591 ns/op	518208648 B/op	19789818 allocs/op
BenchmarkFast-8           	     446	   2366057 ns/op	  943397 B/op	    5471 allocs/op
BenchmarkFastUnfolded-8   	      31	  37762411 ns/op	18051713 B/op	   27238 allocs/op
PASS
ok  	github.com/danfarino/examples/advent-of-code-2023/day12	6.040s
```
