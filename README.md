[![GoDoc](https://godoc.org/github.com/xaionaro-go/weightedshuffle?status.svg)](https://pkg.go.dev/github.com/xaionaro-go/weightedshuffle?tab=doc)

# Shuffle, but with position preference

This package allows to (semi-)randomly order a slice with weight-based preference to get closer to the beginning of the slice.

For example, slice `[1, 2, 3, 4]` with weights `[0, 1, 9, 10]` most likely with result into `[4, 3, 2, 1]` or `[3, 4, 2, 1]`:
* `1` is always in the end because it has zero-weight,
* `2` most likely goes after `3` and `4` because it has much lower weight
* while `4` and `3` have almost the same weight so they will have quite random order relatively to each other.

# Performance
```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/weightedshuffle
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkShuffle/0/Shuffle-12         	 6839139	       180.2 ns/op	     120 B/op	       3 allocs/op
BenchmarkShuffle/0/ShuffleInplace-12  	 8649206	       139.8 ns/op	      56 B/op	       2 allocs/op
BenchmarkShuffle/1/Shuffle-12         	 4564108	       262.4 ns/op	     144 B/op	       4 allocs/op
BenchmarkShuffle/1/ShuffleInplace-12  	11854474	        87.38 ns/op	      56 B/op	       2 allocs/op
BenchmarkShuffle/10/Shuffle-12        	 1289701	       954.4 ns/op	     424 B/op	      10 allocs/op
BenchmarkShuffle/10/ShuffleInplace-12 	 1888671	       627.1 ns/op	     120 B/op	       5 allocs/op
BenchmarkShuffle/100/Shuffle-12       	  114979	      9761 ns/op	    2216 B/op	      13 allocs/op
BenchmarkShuffle/100/ShuffleInplace-12         	  127245	      9563 ns/op	     120 B/op	       5 allocs/op
BenchmarkShuffle/1000/Shuffle-12               	    7822	    132306 ns/op	   16552 B/op	      16 allocs/op
BenchmarkShuffle/1000/ShuffleInplace-12        	    2047	    613564 ns/op	     120 B/op	       5 allocs/op
BenchmarkShuffle/10000/Shuffle-12              	     688	   1745520 ns/op	  386473 B/op	      25 allocs/op
BenchmarkShuffle/10000/ShuffleInplace-12       	      19	  59876093 ns/op	     120 B/op	       5 allocs/op
BenchmarkShuffle/100000/Shuffle-12             	      60	  23312442 ns/op	 4654532 B/op	      35 allocs/op
BenchmarkShuffle/1000000/Shuffle-12            	       4	 282116172 ns/op	45188534 B/op	      45 allocs/op
PASS
ok  	github.com/xaionaro-go/weightedshuffle	20.833s
```
