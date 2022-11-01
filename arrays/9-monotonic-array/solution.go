package main

import "fmt"

// O(n)time | O(1)space
func IsMonotonic(array []int) bool {
	nonDecreasing, nonIncreasing := true, true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			nonDecreasing = false
		}

		if array[i] > array[i-1] {
			nonIncreasing = false
		}
	}

	return nonDecreasing || nonIncreasing
}

func main() {
	array := []int{1, 2, 0}
	response := IsMonotonic(array)
	fmt.Println(response)
}
