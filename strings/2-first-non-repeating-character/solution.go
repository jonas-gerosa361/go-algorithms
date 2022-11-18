package main

import "fmt"

// O(n)time e.g 2 linear but not linear inside linear | O(1)space
func FirstNonRepeatingCharacter(str string) int {
	// Create a hash table to populate with each character number of ocurrences in the string.
	ocurrences := map[rune]int{}
	for _, char := range str {
		ocurrences[char] += 1
	}

	// Iterate over the str again but this time looking for a character with score of 1.
	for index, char := range str {
		if ocurrences[char] == 1 {
			return int(index)
		}
	}

	// Returns -1 if not found.
	return -1
}

func main() {
	str := "faadabcbbebdf"
	response := FirstNonRepeatingCharacter(str)
	fmt.Println(response)
}
