package slicefind

// SliceFind returns the first element in the slice that satisfies the predicate f.
func SliceFind[T any](data []T, f func(T) bool) (T, bool) {
	for _, v := range data {
		if f(v) {
			return v, true
		}
	}
	var result T
	return result, false
}
