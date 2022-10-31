package main

import (
	"fmt"
)

// O(n)time | O(1)space
func MoveElementToEnd(array []int, toMove int) []int {
	leftPointer, rightPointer := 0, len(array)-1

	for leftPointer < rightPointer {
		left, right := array[leftPointer], array[rightPointer]
		if right == toMove {
			rightPointer--
			continue
		}

		if left == toMove {
			array[leftPointer], array[rightPointer] = array[rightPointer], array[leftPointer]
			leftPointer++
			rightPointer--
			continue
		}

		leftPointer++
	}
	return array
}

func main() {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	response := MoveElementToEnd(array, toMove)
	fmt.Println(response)
}
