package main

import "fmt"

// O(n)time of the main array | O(1)space
func IsValidSubsequence(array []int, sequence []int) bool {
	// Create a variable to hold the index of the sequence index - it will start with 0.
	sequenceIndex := 0

	// Iterate in the main array
	for _, arrayNumber := range array {
		// If number is found increment sequence index.
		if arrayNumber == sequence[sequenceIndex] {
			sequenceIndex++
		}

		// If we found all the sequence numbers we can return true.
		if sequenceIndex > len(sequence)-1 {
			return true
		}
	}

	return false
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{26}
	response := IsValidSubsequence(array, sequence)
	fmt.Println(response)
}
