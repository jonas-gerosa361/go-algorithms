package main

import (
	"fmt"
)

// O(n)time | O(1)space
func IsPalindrome(str string) bool {
	leftPointer, rightPointer := 0, len(str)-1
	response := true

	for leftPointer < rightPointer {
		left, right := string(str[leftPointer]), string(str[rightPointer])
		response = false
		if left == right {
			response = true
		}

		leftPointer++
		rightPointer--
	}

	return response
}

func main() {
	word := "tenet"
	response := IsPalindrome(word)
	fmt.Println(response)
}
