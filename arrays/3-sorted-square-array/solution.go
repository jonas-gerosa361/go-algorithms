package main

import (
	"fmt"
	"sort"
)

func SortedSquaredArray(array []int) []int {
	var response []int
	for _, number := range array {
		response = append(response, number*number)
	}

	sort.Ints(response)
	return response
}

func main() {
	array := []int{1, 2, 3, 5, 6, 8, 9}
	response := SortedSquaredArray(array)

	fmt.Println(response)
}
