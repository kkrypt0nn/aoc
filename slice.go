package aoc

import "sort"

// SliceContainsAll checks whether all elements of slice B are in slice A.
func SliceContainsAll[T comparable](sliceA []T, sliceB []T) bool {
	for _, item := range sliceB {
		if !SliceContains(sliceA, item) {
			return false
		}
	}
	return true
}

// SliceContains checks whether an element of type T is contained in the slice.
func SliceContains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// SliceMax returns the maximum value in the slice as an integer.
func SliceMax(slice []int) int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice[0]
}

// SliceMultipleMax returns an amount of maximum values in the slice as a new slice.
func SliceMultipleMax(slice []int, amount int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	var maxValues []int
	for i := 0; i < amount; i++ {
		maxValues = append(maxValues, slice[i])
	}
	return maxValues
}

// SliceMin returns the minimum value in the slice as an integer.
func SliceMin(slice []int) int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice[0]
}

// SliceMultipleMin returns an amount of minimum values in the slice as a new slice.
func SliceMultipleMin(slice []int, amount int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	var minValues []int
	for i := 0; i < amount; i++ {
		minValues = append(minValues, slice[i])
	}
	return minValues
}

// SliceSum returns the sum of the values in the slice.
func SliceSum(slice []int) int {
	sum := 0
	for _, x := range slice {
		sum += x
	}
	return sum
}

// SliceReverse reverses a slice of elements of type T.
func SliceReverse[T any](slice []T) []T {
	var newSlice []T
	for i := len(slice) - 1; i >= 0; i-- {
		newSlice = append(newSlice, slice[i])
	}
	return newSlice
}
