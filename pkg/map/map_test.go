package slice_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func addLetter(str string) string {
	return str + "a"
}

func addNumber(num int) int {
	return num + 1
}

func TestSliceMapWithString(t *testing.T) {
	strs := []string{"a", "b", "c"}
	out := SliceMap(strs, addLetter)
	assert.Equal(t, []string{"aa", "ba", "ca"}, out, "output should be equal")
}

func TestSlickMapWithNumber(t *testing.T) {
	nums := []int{1, 2, 3}
	out := SliceMap(nums, addNumber)
	assert.Equal(t, []int{2, 3, 4}, out, "output should be equal")
}
