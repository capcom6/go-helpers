package slices

// GroupBy groups the elements of a slice based on the result of a key function.
//
// slice: The input slice of type T.
// keyFunc: The key function that maps elements of type T to keys of type K.
// Returns a map where the keys are of type K and the values are slices of type T.
func GroupBy[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	// Handle nil slice edge case
	if slice == nil {
		return make(map[K][]T)
	}

	// Pre-allocate the map for performance optimization
	groups := make(map[K][]T)
	for _, item := range slice {
		key := keyFunc(item)
		groups[key] = append(groups[key], item)
	}
	return groups
}
