package foreach

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var val = 0
var word = ""

func addNumber(num int) {
	val = num + val
	return
}

func addLetter(str string) {
	word = word + str
	return
}

func TestSliceForEachWithNumber(t *testing.T) {
	nums := []int{1, 2, 3}
	SliceForEach(nums, addNumber)
	assert.Equal(t, 6, val, "output should be equal")
}

func TestSliceForEachWithWord(t *testing.T) {
	letters := []string{"a", "b", "c"}
	SliceForEach(letters, addLetter)
	assert.Equal(t, "abc", word, "output should be equal")
}
