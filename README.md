[![GoDoc](https://godoc.org/github.com/xaionaro-go/weightedrandsort?status.svg)](https://pkg.go.dev/github.com/xaionaro-go/weightedrandsort?tab=doc)

# Sort, but with random factor

This package allows to (semi-)randomly order a slice with weight-based preference to get closer to the beginning of the slice.

For example, slice `[1, 2, 3, 4]` with weights `[0, 1, 9, 10]` most likely with result into `[4, 3, 2, 1]` or `[3, 4, 2, 1]`:
* `1` is always in the end because it has zero-weight,
* `2` most likely goes after `3` and `4` because it has much lower weight
* while `4` and `3` have almost the same weight so they will have quite random order relatively to each other.

# Performance
```
goos: linux
goarch: amd64
pkg: github.com/xaionaro-go/weightedrandsort
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkReoder/0/Reorder-12         	 1669608	       742.4 ns/op	     120 B/op	       3 allocs/op
BenchmarkReoder/0/ReorderInplace-12  	 4466482	       318.7 ns/op	      56 B/op	       2 allocs/op
BenchmarkReoder/1/Reorder-12         	 1674543	       765.0 ns/op	     128 B/op	       4 allocs/op
BenchmarkReoder/1/ReorderInplace-12  	 4171033	       349.6 ns/op	      56 B/op	       2 allocs/op
BenchmarkReoder/10/Reorder-12        	 1000000	      2049 ns/op	     232 B/op	       5 allocs/op
BenchmarkReoder/10/ReorderInplace-12 	 1000000	      1343 ns/op	      88 B/op	       3 allocs/op
BenchmarkReoder/100/Reorder-12       	   67473	     19072 ns/op	    1048 B/op	       5 allocs/op
BenchmarkReoder/100/ReorderInplace-12         	   65572	     17665 ns/op	      88 B/op	       3 allocs/op
BenchmarkReoder/1000/Reorder-12               	    5847	    203860 ns/op	    8344 B/op	       5 allocs/op
BenchmarkReoder/1000/ReorderInplace-12        	    1018	   1272364 ns/op	      88 B/op	       3 allocs/op
BenchmarkReoder/10000/Reorder-12              	     409	   3303214 ns/op	   82072 B/op	       5 allocs/op
BenchmarkReoder/10000/ReorderInplace-12       	       9	 116600784 ns/op	      88 B/op	       3 allocs/op
BenchmarkReoder/100000/Reorder-12             	      31	  38014049 ns/op	  802968 B/op	       5 allocs/op
BenchmarkReoder/1000000/Reorder-12            	       3	 425498848 ns/op	 8003736 B/op	       5 allocs/op
PASS
ok  	github.com/xaionaro-go/weightedrandsort	24.772s
```
