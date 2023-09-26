package aoc

import "testing"

func TestSliceContainsAll(t *testing.T) {
	sliceA := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	sliceB := []int{3, 7, 1}
	containsAll := SliceContainsAll(sliceA, sliceB)
	if containsAll {
		t.Log("Successfully checked if all the elements of slice B are in slice A")
		return
	}
	t.Errorf("Failed checking if all the elements of slice B are in slice A (got: %t, expected: true)", containsAll)
}

func TestContains(t *testing.T) {
	theSlice := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	element := 8
	contains := SliceContains(theSlice, element)
	if contains {
		t.Log("Successfully checked if the element is contained in the slice")
		return
	}
	t.Errorf("Failed checking if the element is contained in the slice (got: %t, expected: true)", contains)
}

func TestSliceMax(t *testing.T) {
	theSlice := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	sliceMax := SliceMax(theSlice)
	expected := 10
	if sliceMax == expected {
		t.Log("Successfully got the biggest element in the slice")
		return
	}
	t.Errorf("Failed getting the biggest element in the slice (got: %d, expected: %d)", sliceMax, expected)
}

func TestSliceMultipleMax(t *testing.T) {
	theSlice := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	sliceMax := SliceMultipleMax(theSlice, 2)
	expected := []int{9, 10}
	if SliceContainsAll(sliceMax, expected) && len(sliceMax) == len(expected) {
		t.Log("Successfully got the two biggest numbers in the slice")
		return
	}
	t.Errorf("Failed getting the two biggest numbers in the slice (got: %#v, expected: %#v)", sliceMax, expected)
}

func TestSliceMin(t *testing.T) {
	theSlice := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	sliceMin := SliceMin(theSlice)
	expected := -1
	if sliceMin == expected {
		t.Log("Successfully got the smallest element in the slice")
		return
	}
	t.Errorf("Failed getting the smallest element in the slice (got: %d, expected: %d)", sliceMin, expected)
}

func TestSliceMultipleMin(t *testing.T) {
	theSlice := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	sliceMin := SliceMultipleMin(theSlice, 2)
	expected := []int{-1, 1}
	if SliceContainsAll(sliceMin, expected) && len(sliceMin) == len(expected) {
		t.Log("Successfully got the two smallest numbers in the slice")
		return
	}
	t.Errorf("Failed getting the two smallest numbers in the slice (got: %#v, expected: %#v)", sliceMin, expected)
}

func TestSliceSum(t *testing.T) {
	theSlice := []int{7, 5, 3, 9, 2, 6, 10, 8, 4, 1, -1}
	sliceSum := SliceSum(theSlice)
	expected := 54
	if sliceSum == expected {
		t.Log("Successfully got the sum of the slice")
		return
	}
	t.Errorf("Failed getting the sum of the slice (got: %d, expected: %d)", sliceSum, expected)
}

func TestSliceReverse(t *testing.T) {
	theSlice := []int{7, 5, 3}
	sliceReverse := SliceReverse(theSlice)
	elementA := 3
	elementB := 5
	elementC := 7
	if sliceReverse[0] == elementA && sliceReverse[1] == elementB && sliceReverse[2] == elementC {
		t.Log("Successfully reversed the slice")
		return
	}
	t.Errorf("Failed reversing the slice (got: %#v, expected: %#v)", sliceReverse, []int{elementA, elementB, elementC})
}
