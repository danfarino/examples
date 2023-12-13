```
$ ./run.sh
goos: linux
goarch: amd64
pkg: github.com/danfarino/examples/advent-of-code-2023/day12
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkSlow-8           	       1	1222186637 ns/op	518208992 B/op	19789819 allocs/op
BenchmarkFast-8           	     481	   2387531 ns/op	  943435 B/op	    5471 allocs/op
BenchmarkFastUnfolded-8   	      30	  39446518 ns/op	18053208 B/op	   27251 allocs/op
PASS
ok  	github.com/danfarino/examples/advent-of-code-2023/day12	6.134s
```
