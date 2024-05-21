package slices

// Map applies the given function to each element of the input slice and returns a new slice with the results.
//
// slice: The input slice of type T.
// f: The function that maps elements of type T to elements of type U.
// Returns a slice with values of type U.
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
