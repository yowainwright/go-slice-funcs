package slicemap

// SliceMap Mapping Function for Generic types
func SliceMap[T any](data []T, f func(T) T) []T {
	result := make([]T, len(data))
	for i, v := range data {
		result[i] = f(v)
	}
	return result
}
