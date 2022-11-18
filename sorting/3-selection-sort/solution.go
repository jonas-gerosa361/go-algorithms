package main

import "fmt"

// O(n²)time | O(1)space
func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				swap(i, j, array)
			}
		}
	}
	return array
}

func swap(i int, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	response := SelectionSort(array)
	fmt.Println(response)
}
