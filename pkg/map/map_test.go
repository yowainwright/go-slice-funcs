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

func updateObject(obj map[string]string) map[string]string {
	obj["name"] = "Jeff"
	var newObj = obj
	newObj["fullName"] = obj["name"] + " " + obj["lastName"]
	return newObj
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

func TestSliceMapWithObject(t *testing.T) {
	objects := []map[string]string{
		{"name": "John", "lastName": "Doe"},
		{"name": "Jane", "lastName": "Doe"},
	}
	out := SliceMap(objects, updateObject)
	assert.Equal(t, []map[string]string{
		{"name": "Jeff", "lastName": "Doe", "fullName": "Jeff Doe"},
		{"name": "Jeff", "lastName": "Doe", "fullName": "Jeff Doe"},
	}, out, "output should be equal")
}
