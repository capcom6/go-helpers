package maps

// MapValues applies the given function to each value of the input map and returns
// a new map with the same keys and the results of the function as values.
func MapValues[K comparable, V any, R any](m map[K]V, f func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}
