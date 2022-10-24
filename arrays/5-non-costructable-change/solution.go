package main

import (
	"fmt"
	"sort"
)

// O(nlogn) time | O(1) space
func NonConstructibleChange(coins []int) int {
	//Start sorting the array.
	sort.Ints(coins)
	minimunChange := 0

	for index := 0; index <= len(coins)-1; index++ {
		if coins[index] > minimunChange+1 {
			return minimunChange + 1
		}

		minimunChange += coins[index]
	}

	return minimunChange + 1
}

func main() {
	array := []int{1, 1}
	result := NonConstructibleChange(array)
	fmt.Println(result)
}
