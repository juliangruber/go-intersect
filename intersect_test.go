package intersect

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSimple(t *testing.T) {
	s := Simple([]int{1}, []int{2})
	assert.Equal(t, s, []interface{}{})

	s = Simple([]int{1, 2}, []int{2})
	assert.Equal(t, s, []interface{}{2})
}
