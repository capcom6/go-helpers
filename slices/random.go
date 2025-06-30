//nolint:gosec // no need to strong randomness
package slices

import (
	"errors"
	"math/rand/v2"
)

var ErrEmptySlice = errors.New("slice is empty")

// Random returns a random element from the given slice.
func Random[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		return *new(T), ErrEmptySlice
	}
	if len(slice) == 1 {
		return slice[0], nil
	}

	return slice[rand.IntN(len(slice))], nil
}
