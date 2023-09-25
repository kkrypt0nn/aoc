package aoc

import (
	"strconv"
	"strings"
)

// Reverse reverses a slice of elements of type T.
func Reverse[T any](slice []T) []T {
	var newSlice []T
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}

// ReverseString reverses a string.
func ReverseString(str string) string {
	var result string
	for _, i := range str {
		result = string(i) + result
	}
	return result
}

// Atoi converts a string to an integer without caring about errors.
func Atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// ContainsSubstr returns true if the string contains the given substring.
func ContainsSubstr(str string, substr string) bool {
	for _, i := range substr {
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}

// ContainsExactly returns true if the string contains the given substring and has the same length.
func ContainsExactly(str string, substr string) bool {
	if len(str) != len(substr) {
		return false
	}
	for _, i := range substr {
		// Yes that is a bit wacky if the string contains the same character multiple times, though worked for the challenge in 2021.
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}
