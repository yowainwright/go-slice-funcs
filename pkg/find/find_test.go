package slicefind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isEven(num int) bool {
	return num%2 == 0
}

func TestSliceFindWithEvenNumberPass(t *testing.T) {
	nums := []int{1, 2, 3}
	out, ok := SliceFind(nums, isEven)
	assert.Equal(t, 2, out, "output should be equal")
	assert.Equal(t, true, ok, "output should be equal")
}

func TestSliceFindWithEvenNumberFail(t *testing.T) {
	nums := []int{1, 3}
	out, ok := SliceFind(nums, isEven)
	assert.Equal(t, 0, out, "output should be equal")
	assert.Equal(t, false, ok, "output should be equal")
}
