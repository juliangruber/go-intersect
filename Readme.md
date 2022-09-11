
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

| **Name**                               |   **Runs** |     **Average** | **Allocated** |     **Allocations from heap** |
|----------------------------------------|-----------:|----------------:|--------------:|------------------------------:|
| BenchmarkHash/Size_1-_interface-8      |  4.631.326 |     248.1 ns/op |       80 B/op |                   5 allocs/op |
| BenchmarkHash/Size_1-_generics-8       | 22.551.091 |     46.19 ns/op |        8 B/op |                   1 allocs/op |
| BenchmarkHash/Size_10-_interface-8     |    581.262 |     2.042 ns/op |      547 B/op |                  24 allocs/op |
| BenchmarkHash/Size_10-_generics-8      |  2.043.096 |     549.6 ns/op |      187 B/op |                   2 allocs/op |
| BenchmarkHash/Size_100-_interface-8    |     59.907 |    18.721 ns/op |    7.325 B/op |                 213 allocs/op |
| BenchmarkHash/Size_100-_generics-8     |    126.622 |     8.221 ns/op |    3.359 B/op |                  19 allocs/op |
| BenchmarkHash/Size_1000-_interface-8   |      7.114 |   22.1721 ns/op |  112.405 B/op |               2.038 allocs/op |
| BenchmarkHash/Size_1000-_generics-8    |     13.345 |    91.197 ns/op |   53.323 B/op |                  74 allocs/op |
| BenchmarkHash/Size_10000-_interface-8  |        523 | 2.067.330 ns/op |  875.731 B/op |              20.173 allocs/op |
| BenchmarkHash/Size_10000-_generics-8   |      1.264 |   828.214 ns/op |  427.541 B/op |                 320 allocs/op |

## Sorted Generic V2

`SortedGenericV2` function allows increase performance for the few times, but can be unsafe in some cases cause will change order for elements in the left array.

| **Name**                                           |    **Runs** |     **Average** | **Allocated** | **Allocations from heap** |
|----------------------------------------------------|------------:|----------------:|--------------:|--------------------------:|
| BenchmarkSortedGeneric/SortedGeneric_-_1-16        |  41 944 108 |     28.25 ns/op |        8 B/op |               1 allocs/op |
| BenchmarkSortedGeneric/SortedGenericV2_-_1-16      | 199 261 370 |     5.992 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGeneric_-_10-16       |  10 445 590 |     111.6 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGenericV2_-_10-16     |  20 496 318 |     57.31 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGeneric_-_100-16      |     745 248 |     1 624 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGenericV2_-_100-16    |   1 909 890 |     626.2 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGeneric_-_1000-16     |      38 613 |    31 097 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGenericV2_-_1000-16   |     188 733 |     6 367 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGeneric_-_10000-16    |       3 799 |   308 313 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGenericV2_-_10000-16  |       9 190 |   134 617 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGeneric_-_100000-16   |         327 | 3 679 350 ns/op |        0 B/op |               0 allocs/op |
| BenchmarkSortedGeneric/SortedGenericV2_-_100000-16 |       1 537 |   783 528 ns/op |        0 B/op |               0 allocs/op |
