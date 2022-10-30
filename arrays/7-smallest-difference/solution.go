package main

import (
	"fmt"
	"math"
	"sort"
)

// O(nlog(n) & nlog(m))time | O(1)space
func SmallestDifference(array1, array2 []int) []int {
	// Start sorting both arrays
	sort.Ints(array1)
	sort.Ints(array2)

	response := []int{}
	leftPointer, rightPointer := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32

	for leftPointer < len(array1) && rightPointer < len(array2) {
		leftValue, rightValue := array1[leftPointer], array2[rightPointer]

		if rightValue == leftValue {
			return []int{leftValue, rightValue}
		}

		if leftValue < rightValue {
			current = rightValue - leftValue
			leftPointer++
		}

		if rightValue < leftValue {
			current = leftValue - rightValue
			rightPointer++
		}

		if smallest > current {
			smallest = current
			response = []int{leftValue, rightValue}
		}
	}

	return response
}

func main() {
	array1 := []int{-1, 5, 10, 20, 28, 3}
	array2 := []int{26, 134, 135, 15, 17}
	response := SmallestDifference(array1, array2)
	fmt.Println(response)
}
