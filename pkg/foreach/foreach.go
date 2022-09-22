package foreach

// SliceForEach iterates over a slice and applies a function to each element
func SliceForEach[T any](data []T, f func(T)) {
	for _, v := range data {
		f(v)
	}
}
