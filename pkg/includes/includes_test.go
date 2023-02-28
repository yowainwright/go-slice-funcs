package includes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIncludes(t *testing.T) {
	nums := []int{1, 2, 3}
	assert.True(t, SliceIncludes(nums, 1), "output should be true")
	assert.False(t, SliceIncludes(nums, 4), "output should be false")
}
