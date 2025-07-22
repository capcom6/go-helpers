package slices_test

import (
	"reflect"
	"testing"

	"github.com/capcom6/go-helpers/slices"
)

func TestGroupBy(t *testing.T) {
	testCases := []struct {
		name     string
		slice    []int
		keyFunc  func(int) int
		expected map[int][]int
	}{
		{
			name:     "empty slice",
			slice:    []int{},
			keyFunc:  func(i int) int { return i % 2 },
			expected: map[int][]int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			keyFunc:  func(i int) int { return i % 2 },
			expected: map[int][]int{},
		},
		{
			name:    "even and odd",
			slice:   []int{1, 2, 3, 4, 5},
			keyFunc: func(i int) int { return i % 2 },
			expected: map[int][]int{
				0: {2, 4},
				1: {1, 3, 5},
			},
		},
		{
			name:    "all even",
			slice:   []int{2, 4, 6, 8},
			keyFunc: func(i int) int { return i % 2 },
			expected: map[int][]int{
				0: {2, 4, 6, 8},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := slices.GroupBy(tc.slice, tc.keyFunc)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("GroupBy(%v) = %v, expected %v", tc.slice, result, tc.expected)
			}
		})
	}
}

func BenchmarkGroupBy(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	keyFunc := func(i int) int { return i % 2 }
	for range b.N {
		slices.GroupBy(slice, keyFunc)
	}
}
