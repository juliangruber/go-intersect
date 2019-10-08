package intersect

import (
	"testing"

	"github.com/bmizerany/assert"
)

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
