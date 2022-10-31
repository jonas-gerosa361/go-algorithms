package main

import "fmt"

// O(n) space | O(n^2) time
func InsertionSort(array []int) []int {
	for index := 1; index < len(array); index++ {
		sortedIndex := index
		for sortedIndex > 0 && array[sortedIndex] < array[sortedIndex-1] {
			swap(sortedIndex, sortedIndex-1, array)
			sortedIndex--
		}
	}
	return array
}

// Helper function to swap 2 values within an array.
func swap(i int, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	response := InsertionSort(array)
	fmt.Println(response)
}
