package slices_test

import (
	"testing"

	"github.com/capcom6/go-helpers/slices"
)

func TestRandom(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
	}{
		{"empty slice", []int{}},
		{"slice with one element", []int{1}},
		{"slice with multiple elements", []int{1, 2, 3, 4, 5}},
		{"slice with duplicate elements", []int{1, 2, 2, 3, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.slice) == 0 {
				// Test that Random returns an error on empty slice
				_, err := slices.Random(tt.slice)
				if err == nil {
					t.Errorf("Random did not return an error on empty slice")
				}
			} else {
				// Test that Random returns an element from the slice
				result, err := slices.Random(tt.slice)
				if err != nil {
					t.Errorf("Random returned an error: %v", err)
				}
				if !contains(tt.slice, result) {
					t.Errorf("Random returned %v, which is not in the slice %v", result, tt.slice)
				}
			}
		})
	}
}

func contains(slice []int, elem int) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}
	return false
}
