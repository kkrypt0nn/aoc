package aoc

import "testing"

func TestMapSum(t *testing.T) {
	theMap := map[int]int{1: 2, 2: 3, 3: 4, 4: 5, 5: 6, 6: 7, 7: 8, 8: 9, 9: 10}
	mapSum := MapSum(theMap)
	expected := 54
	if mapSum == expected {
		t.Log("Successfully got the sum of the map")
		return
	}
	t.Errorf("Failed getting the sum of the map (got: %d, expected: %d)", mapSum, expected)
}
