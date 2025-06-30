package slices_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/capcom6/go-helpers/slices"
)

var (
	errOnFirst = errors.New("error on first")
	errOnMid   = errors.New("error on mid")
	errOnLast  = errors.New("error on last")
)

func TestMap(t *testing.T) {
	// Test when the input slice is empty
	t.Run("Empty Slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := slices.Map(input, func(x int) int { return x })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has positive integers
	t.Run("Positive Integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 4, 9, 16, 25}
		result := slices.Map(input, func(x int) int { return x * x })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has negative integers
	t.Run("Negative Integers", func(t *testing.T) {
		input := []int{-1, -2, -3, -4, -5}
		expected := []int{1, 4, 9, 16, 25}
		result := slices.Map(input, func(x int) int { return x * x })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has mixed integers
	t.Run("Mixed Integers", func(t *testing.T) {
		input := []int{-1, 0, 1, -2, 2}
		expected := []int{1, 0, 1, 2, 2}
		result := slices.Map(input, func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		})
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has strings
	t.Run("Strings", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		expected := []string{"APPLE", "BANANA", "CHERRY"}
		result := slices.Map(input, strings.ToUpper)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func testMapOrErrorSuccess[T any](t *testing.T, input []T, transformFunc func(T) (T, error), expected []T) {
	result, err := slices.MapOrError(input, transformFunc)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func testMapOrErrorError[T any](t *testing.T, input []T, transformFunc func(T) (T, error), expectedErr error) {
	_, err := slices.MapOrError(input, transformFunc)
	if err == nil || !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}

func TestMapOrErrorEmptySlice(t *testing.T) {
	t.Run("Empty Slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		testMapOrErrorSuccess(t, input, func(x int) (int, error) {
			return x, nil
		}, expected)
	})
}

func TestMapOrErrorSuccessfulTransformations(t *testing.T) {
	t.Run("Successful Transformations - Integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 4, 9, 16, 25}
		testMapOrErrorSuccess(t, input, func(x int) (int, error) {
			return x * x, nil
		}, expected)
	})

	t.Run("String Transformations", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		expected := []string{"APPLE", "BANANA", "CHERRY"}
		testMapOrErrorSuccess(t, input, func(s string) (string, error) {
			return strings.ToUpper(s), nil
		}, expected)
	})
}

func TestMapOrErrorErrorCases(t *testing.T) {
	t.Run("Error in First Element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		testMapOrErrorError(t, input, func(x int) (int, error) {
			if x == 1 {
				return 0, errOnFirst
			}
			return x * x, nil
		}, errOnFirst)
	})

	t.Run("Error in Middle Element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		testMapOrErrorError(t, input, func(x int) (int, error) {
			if x == 3 {
				return 0, errOnMid
			}
			return x * x, nil
		}, errOnMid)
	})

	t.Run("Error in Last Element", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		testMapOrErrorError(t, input, func(x int) (int, error) {
			if x == 5 {
				return 0, errOnLast
			}
			return x * x, nil
		}, errOnLast)
	})

	t.Run("String Transformation Error", func(t *testing.T) {
		input := []string{"apple", "banana", "cherry"}
		testMapOrErrorError(t, input, func(s string) (string, error) {
			if s == "banana" {
				return "", errOnMid
			}
			return strings.ToUpper(s), nil
		}, errOnMid)
	})
}
