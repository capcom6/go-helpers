package slices

import (
	"errors"
	"math/rand"
)

// Random returns a random element from the given slice.
func Random[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		return *new(T), errors.New("slice is empty")
	}
	if len(slice) == 1 {
		return slice[0], nil
	}

	return slice[rand.Intn(len(slice))], nil
}
