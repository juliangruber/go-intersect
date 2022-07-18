
# go-intersect

  Find the intersection of two iterable values.

  This library provides multiple implementations which each have their strong and weak points.

  Read the [docs](http://godoc.org/github.com/juliangruber/go-intersect).

  [![Build Status](https://travis-ci.com/juliangruber/go-intersect.svg?branch=master)](https://travis-ci.com/juliangruber/go-intersect)

## Installation

```bash
$ go get github.com/juliangruber/go-intersect
```

## Example

```go
import "github.com/juliangruber/go-intersect"
import "fmt"

func main() {
  a := []int{1, 2, 3}
  b := []int{2, 3, 4}
  fmt.Println(intersect.Simple(a, b))
}
```

## Generics

Go v1.18 now supports generics which increases performance compared to reflection. Below is a benchmark results between reflection and generics of the hash intersection functions. Slice sizes of 1, 10, 100 1.000 and 10.000 are used in benchmark.

The average time and allocated bytes are more than halved when using generics, and the number of times bytes has to be allocated from the heap are extremely improved.

| **Name**                              	| **Runs** 	| **Average** 	| **Allocated** 	| **Allocations from heap** 	|
|---------------------------------------	|-----------------:	|--------------------:	|---------------------:	|-------------------------------------------:	|
| BenchmarkHash/Size_1-_interface-8     	| 4.631.326         	| 248.1 ns/op        	| 80 B/op             	| 5 allocs/op                               	|
| BenchmarkHash/Size_1-_generics-8      	| 22.551.091        	| 46.19 ns/op        	| 8 B/op              	| 1 allocs/op                               	|
| BenchmarkHash/Size_10-_interface-8    	| 581.262          	| 2.042 ns/op         	| 547 B/op            	| 24 allocs/op                              	|
| BenchmarkHash/Size_10-_generics-8     	| 2.043.096         	| 549.6 ns/op        	| 187 B/op            	| 2 allocs/op                               	|
| BenchmarkHash/Size_100-_interface-8   	| 59.907           	| 18.721 ns/op        	| 7.325 B/op           	| 213 allocs/op                             	|
| BenchmarkHash/Size_100-_generics-8    	| 126.622          	| 8.221 ns/op         	| 3.359 B/op           	| 19 allocs/op                              	|
| BenchmarkHash/Size_1000-_interface-8  	| 7.114            	| 22.1721 ns/op       	| 112.405 B/op         	| 2.038 allocs/op                            	|
| BenchmarkHash/Size_1000-_generics-8   	| 13.345           	| 91.197 ns/op        	| 53.323 B/op          	| 74 allocs/op                              	|
| BenchmarkHash/Size_10000-_interface-8 	| 523             	| 2.067.330 ns/op      	| 875.731 B/op         	| 20.173 allocs/op                           	|
| BenchmarkHash/Size_10000-_generics-8  	| 1.264            	| 828.214 ns/op       	| 427.541 B/op         	| 320 allocs/op                             	|

