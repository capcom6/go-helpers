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

	return slice[rand.Intn(len(slice))], nil
}
