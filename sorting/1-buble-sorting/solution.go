package main

import "fmt"

// Best O(n) time | O(1)space
// Average O(n^2) time | O(1)space
// Worst O(n^2) time | O(1)space
func BubbleSort(array []int) []int {
	sorted := false
	// Used to skip the last sorted number on each iteration.
	counter := 0
	for sorted != true {
		sorted = true
		for index := 0; index < len(array)-1-counter; index++ {
			if array[index] > array[index+1] {
				swap(index, index+1, array)
				sorted = false
			}
		}
		counter++
	}

	return array
}

func swap(index int, nextIndex int, array []int) {
	array[index], array[nextIndex] = array[nextIndex], array[index]
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	BubbleSort(array)

	fmt.Println(array)
}
