package main

import (
	"fmt"
	"sort"
)

// O(n^2)time | O(n)space
func ThreeNumberSum(array []int, target int) [][]int {
	// Start sorting the array.
	sort.Ints(array)
	response := [][]int{}

	// Go left and right searching for:
	for index, number := range array {
		left, right := index+1, len(array)-1
		for left < right {
			// leftNumber + nextLeftnumber + rightNumber == target.
			currentSum := number + array[left] + array[right]
			if currentSum == target {
				response = append(response, []int{number, array[left], array[right]})
				// Move both pointers.
				left++
				right--
			}

			if currentSum < target {
				// Move only left pointer.
				left++
			}

			if currentSum > target {
				// Move only right pointer.
				right--
			}
		}
	}

	return response
}

func main() {
	array := []int{1, 2, 3}
	targetSum := 7
	response := ThreeNumberSum(array, targetSum)
	fmt.Println(response)
}
