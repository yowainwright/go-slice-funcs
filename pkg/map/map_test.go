package slice_map

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func addLetter(str string) string {
	return str + "a"
}

func TestSliceMapWithString(t *testing.T) {
	strs := []string{"a", "b", "c"}
	out := SliceMap(strs, addLetter)
	assert.Equal(t, []string{"aa", "ba", "ca"}, out, "output should be equal")
}
