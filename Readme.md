
# go-intersect

  Find the intersection of two iterable values.

  This library provides multiple implementations which each have their strong and weak points.

  Read the [docs](https://pkg.go.dev/github.com/juliangruber/go-intersect/v2).

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
  fmt.Println(intersect.SimpleGeneric(a, b))
}
```

## Generics

Go v1.18 now supports generics which increases performance compared to reflection. Below are benchmark results between reflection and generics of the hash intersection functions. Slice sizes of 1, 10, 100 1.000 and 10.000 are used.

The average time and allocated bytes are more than halved when using generics, and the number of times bytes have to be allocated from the heap are improved significantly.

| **Name**                               |   **Runs** |     **Average** | **Allocated** | **Allocations from heap** |
|----------------------------------------|-----------:|----------------:|--------------:|--------------------------:|
| BenchmarkHash/Size_1-_interface-16     |  5 845 340 |     202.9 ns/op |       80 B/op |               5 allocs/op |
| BenchmarkHash/Size_1-_generics-16      | 29 534 389 |     39.86 ns/op |        8 B/op |               1 allocs/op |
| BenchmarkHash/Size_10-_interface-16    |    800 064 |     1 438 ns/op |      515 B/op |              24 allocs/op |
| BenchmarkHash/Size_10-_generics-16     |  2 459 796 |     482.6 ns/op |      170 B/op |               2 allocs/op |
| BenchmarkHash/Size_100-_interface-16   |     81 973 |    14 782 ns/op |    6 907 B/op |             212 allocs/op |
| BenchmarkHash/Size_100-_generics-16    |    184 526 |     6 467 ns/op |    3 052 B/op |              19 allocs/op |
| BenchmarkHash/Size_1000-_interface-16  |      7 489 |   164 909 ns/op |  101 369 B/op |           2 036 allocs/op |
| BenchmarkHash/Size_1000-_generics-16   |     16 734 |    70 601 ns/op |   47 743 B/op |              66 allocs/op |
| BenchmarkHash/Size_10000-_interface-16 |        754 | 1 592 308 ns/op |  836 264 B/op |          20 216 allocs/op |
| BenchmarkHash/Size_10000-_generics-16  |      1 802 |   652 946 ns/op |  389 466 B/op |             255 allocs/op |


