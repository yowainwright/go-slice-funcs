package slicefilter

// SliceFilter Function for Generic types
func SliceFilter[T any](data []T, f func(T) bool) []T {
	result := make([]T, 0, len(data))
	for _, v := range data {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
