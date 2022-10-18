/*
QUESTION:
Writes a function that takes a non empty array with integers and a integer with the target sum.
There will always be two numbers sum which will reache the target sum.

Discover the two numbers which summed are equals to the target value.
*/

package main

import (
	"fmt"
	"sort"
)

// O(n)time | O(n)space
func TwoNumberSum(array []int, target int) []int {
	// Sort array first.
	sort.Ints(array)
	// With array sorted we can go left and right to find the sum
	left, right := 0, len(array)-1
	// Iterate in the array while left < right.
	for left < right { // linear in O(time)
		currentSum := array[left] + array[right] // linear in O(space)
		if currentSum == target {
			return []int{array[left], array[right]}
		}

		if currentSum < target {
			left += 1
		}

		if currentSum > target {
			right -= 1

		}
	}

	// return empty array if not found.
	return []int{}
}

// Func used only for testing.
func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	result := TwoNumberSum(array, target)

	fmt.Println(result)
}
