package main

import "fmt"

func BinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	response := -1
	indexInArray := left+right/2 < len(array)-1
	for left < right && indexInArray {
		index := left + right/2
		if array[index] == target {
			response = index
			break
		}

		if array[index] > target {
			right--
		}

		if array[index] < target {
			left++
		}
	}

	return response
}

func main() {
	array := []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}
	target := 61
	response := BinarySearch(array, target)
	fmt.Println(response)
}
