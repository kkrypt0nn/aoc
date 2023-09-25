package aoc

// MapSum returns the sum of the values in the map.
func MapSum(theMap map[int]int) int {
	sum := 0
	for _, x := range theMap {
		sum += x
	}
	return sum
}
