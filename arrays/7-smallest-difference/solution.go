package main

import (
	"fmt"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	// Start sorting both arrays
	sort.Ints(array1)
	sort.Ints(array2)

	response := []int{}
	leftPointer, rightPointer := 0, 0
	minFound := false

	minChange := 0
	for minFound != true {
		currentChange := array1[leftPointer] - array2[rightPointer]
		fmt.Println(currentChange)
		if currentChange < 0 {
			minChange = currentChange
			response = []int{array1[leftPointer], array2[rightPointer]}
		}

		if array1[leftPointer] < array2[rightPointer] && leftPointer < len(array1)-1 {
			leftPointer++
		}

		if array1[leftPointer] > array2[rightPointer] && rightPointer < len(array2)-1 {
			rightPointer++
		}

		fmt.Println(minChange)
	}

	return response
}

func main() {
	array1 := []int{-1, 5, 10, 20, 28, 3}
	array2 := []int{26, 134, 135, 15, 17}
	response := SmallestDifference(array1, array2)
	fmt.Println(response)
}
