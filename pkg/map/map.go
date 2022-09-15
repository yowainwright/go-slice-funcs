package map

import "fmt"

// A Slice Mapping Function for Generic types
func Map (data []T, f func(T) T) []T {
		result := make([]T, len(data))
		for i, v := range data {
				result[i] = f(v)
		}
		return result
}
