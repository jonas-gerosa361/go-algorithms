package main

import (
	"fmt"
	"sort"
)

// TODO: Implement better answer, try to remove the sort method, we can do this keeping track of the minimun and max value  on response.
func SortedSquaredArray(array []int) []int {
	var response []int
	left, right := 0, len(array)-1

	for left <= right {
		if left == right {
			response = append(response, array[left]*array[left])
			break
		}

		response = append(response, array[left]*array[left])
		response = append(response, array[right]*array[right])

		left++
		right--
	}

	sort.Ints(response)
	return response
}

func main() {
	array := []int{1, 2, 3, 5, 6, 8, 9}
	response := SortedSquaredArray(array)

	fmt.Println(response)
}
