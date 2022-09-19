package slicereduce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestSliceReduce(t *testing.T) {
	nums := []int{1, 2, 3}
	out := SliceReduce(nums, add)
	assert.Equal(t, 6, out, "output should be equal")
}
