package intersect

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/bmizerany/assert"
)

func TestSimpleGeneric(t *testing.T) {
	s := SimpleGeneric([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []int{})

	s = SimpleGeneric([]int{1, 2}, []int{2})
	assert.Equal(t, s, []int{2})
}

func TestSortedGeneric(t *testing.T) {
	s := SortedGeneric([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []int{})

	s = SortedGeneric([]int{1, 2}, []int{2})
	assert.Equal(t, s, []int{2})
}

func TestHashGeneric(t *testing.T) {
	s := HashGeneric([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []int{})

	s = HashGeneric([]int{1, 2}, []int{2})
	assert.Equal(t, s, []int{2})
}

func TestSimple(t *testing.T) {
	s := Simple([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []interface{}{})

	s = Simple([]int{1, 2}, []int{2})
	assert.Equal(t, s, []interface{}{2})
}

func TestSorted(t *testing.T) {
	s := Sorted([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []interface{}{})

	s = Sorted([]int{1, 2}, []int{2})
	assert.Equal(t, s, []interface{}{2})
}

func TestHash(t *testing.T) {
	s := Hash([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []interface{}{})

	s = Hash([]int{1, 2}, []int{2})
	assert.Equal(t, s, []interface{}{2})
}

var blackholeHashGeneric []int
var blackholeHash []interface{}

func BenchmarkHash(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1_000, 10_000} {
		aSlice := createRandomSlice(v)
		bSlice := createRandomSlice(v)
		b.Run(fmt.Sprintf("Size %d- interface", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				blackholeHash = Hash(aSlice, bSlice)
			}
		})
		b.Run(fmt.Sprintf("Size %d- generics", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				blackholeHashGeneric = HashGeneric(aSlice, bSlice)
			}
		})
	}
}

func createRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(int(math.Pow(float64(size), 2)))
	}
	return slice
}
