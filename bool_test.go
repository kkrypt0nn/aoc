package aoc

import "testing"

func TestBoolToInt(t *testing.T) {
	intTrue := BoolToInt(true)
	intFalse := BoolToInt(false)
	expectedTrue := 1
	expectedFalse := 0
	if intTrue == expectedTrue && intFalse == expectedFalse {
		t.Log("Successfully converted a bool value to its int representation")
		return
	}
	t.Errorf("Failed converting a bool value to its int representation (got: true=%d/false=%d, expected: true=%d/false=%d)", intTrue, intFalse, expectedTrue, expectedFalse)
}
