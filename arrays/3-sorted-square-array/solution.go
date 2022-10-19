package main

import (
	"fmt"
)

// O(n) time | O(n) space
func SortedSquaredArray(array []int) []int {
	sortedArray := make([]int, len(array))

	minIndex, maxIndex := 0, len(array)-1

	for index := len(array) - 1; index >= 0; index-- {
		minValue := array[minIndex]
		maxValue := array[maxIndex]

		if abs(minValue) > abs(maxValue) {
			sortedArray[index] = minValue * minValue
			minIndex += 1
		}

		if abs(minValue) <= abs(maxValue) {
			sortedArray[index] = maxValue * maxValue
			maxIndex -= 1
		}
	}

	return sortedArray
}

// Helper function to deal with negative values.
func abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func main() {
	array := []int{1, 2, 3, 5, 6, 8, 9}
	response := SortedSquaredArray(array)

	fmt.Println(response)
}
