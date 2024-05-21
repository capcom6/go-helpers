package slices

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	// Test when the input slice is empty
	t.Run("Empty Slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := Map(input, func(x int) int { return x })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has positive integers
	t.Run("Positive Integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 4, 9, 16, 25}
		result := Map(input, func(x int) int { return x * x })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has negative integers
	t.Run("Negative Integers", func(t *testing.T) {
		input := []int{-1, -2, -3, -4, -5}
		expected := []int{1, 4, 9, 16, 25}
		result := Map(input, func(x int) int { return x * x })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	// Test when the input slice has mixed integers
	t.Run("Mixed Integers", func(t *testing.T) {
		input := []int{-1, 0, 1, -2, 2}
		expected := []int{1, 0, 1, 2, 2}
		result := Map(input, func(x int) int {
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
		result := Map(input, func(s string) string { return strings.ToUpper(s) })
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
