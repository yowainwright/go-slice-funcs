package includes

func SliceIncludes[T any](data []T, value T) bool {
	castedValue := interface{}(value)
	for _, v := range data {
		castedv := interface{}(v)
		if castedv == castedValue {
			return true
		}
	}
	return false
}
