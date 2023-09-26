package aoc

import (
	"strconv"
	"strings"
)

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

// ContainsCharacters returns true if the string contains the given characters anywhere in any order.
func ContainsCharacters(str string, characters string) bool {
	for _, i := range characters {
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}

// ContainsExactly returns true if the string contains the given characters and has the same length.
func ContainsExactly(str string, characters string) bool {
	if len(str) != len(characters) {
		return false
	}
	for _, i := range characters {
		// Yes that is a bit wacky if the string contains the same character multiple times, though worked for the challenge in 2021.
		if !strings.Contains(str, string(i)) {
			return false
		}
	}
	return true
}
