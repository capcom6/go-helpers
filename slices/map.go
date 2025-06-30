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

// MapOrError applies the given function to each element of the input slice,
// collecting results or returning the first error encountered.
//
// slice: The input slice of type T.
// f: The function that maps elements of type T to elements of type U and may return an error.
// Returns a slice with values of type U or an error if one occurred.
func MapOrError[T, U any](slice []T, f func(T) (U, error)) ([]U, error) {
	result := make([]U, len(slice))
	for i, v := range slice {
		u, err := f(v)
		if err != nil {
			return nil, err
		}
		result[i] = u
	}
	return result, nil
}
