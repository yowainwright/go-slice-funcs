package slicereduce

// SliceReduce Reducing Function for Generic types
func SliceReduce[T any](data []T, f func(T, T) T) T {
	result := data[0]
	for _, v := range data[1:] {
		result = f(result, v)
	}
	return result
}
