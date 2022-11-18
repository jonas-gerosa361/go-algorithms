package main

import "fmt"

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
	Depth       int
}

var level int = 0
var depth int = 0

func NodeDepths(root *BinaryTree) int {
	root.Depth = depth
	fmt.Println(level, depth, root.Value)

	if root.Left != nil && root.Right != nil {
		level++
	}

	if root.Left != nil {
		root.Left.Depth = depth + 1
		checkDepth(root.Left)
	}

	if root.Right != nil {
		checkDepth(root.Right)
	}

	return depth
}

func checkDepth(root *BinaryTree) {

}

func main() {
	input := BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}
	response := NodeDepths(&input)
	fmt.Println(response)
}
