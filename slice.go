package aoc

import "sort"

// Contains checks whether an element of type T is contained in the slice or not.
func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}

// Max returns the maximum value in the slice as an integer.
func Max(slice []int) int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	return slice[0]
}

// MaxSlice returns an amount of maximum values in the slice as a new slice.
func MaxSlice(slice []int, amount int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
	var max []int
	for i := 0; i < amount; i++ {
		max = append(max, slice[i])
	}
	return max
}

// Min returns the minimum value in the slice as an integer.
func Min(slice []int) int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice[0]
}

// MinSlice returns an amount of minimum values in the slice as a new slice.
func MinSlice(slice []int, amount int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	var min []int
	for i := 0; i < amount; i++ {
		min = append(min, slice[i])
	}
	return min
}

// Sum returns the sum of the values in the slice.
func Sum(slice []int) int {
	sum := 0
	for _, x := range slice {
		sum += x
	}
	return sum
}
