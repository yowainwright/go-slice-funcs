package slice_filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isEven(num int) bool {
	return num%2 == 0
}

func is4Letters(str string) bool {
	return len(str) == 4
}

func TestSliceFilterWithNumber(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	out := SliceFilter(nums, isEven)
	assert.Equal(t, []int{2, 4}, out, "output should be equal")
}

func TestSliceFilterWithStringsThatAre4Letters(t *testing.T) {
	strs := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "bbbb", "cccc"}
	out := SliceFilter(strs, is4Letters)
	assert.Equal(t, []string{"aaaa", "bbbb", "cccc"}, out, "output should be equal")
}
