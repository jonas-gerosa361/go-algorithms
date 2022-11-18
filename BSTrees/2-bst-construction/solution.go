package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value >= tree.Value && tree.Right == nil {
		tree.Right = new(BST)
		tree.Right.Value = value
	}

	if value < tree.Value && tree.Left == nil {
		tree.Left = new(BST)
		tree.Left.Value = value
	}

	return tree
}

func (tree *BST) Contains(value int) bool {
	if value > tree.Value {
		return FindInTree(tree.Right, value)
	}

	if value < tree.Value {
		return FindInTree(tree.Left, value)
	}

	return false
}

func FindInTree(currentNode *BST, value int) bool {
	if currentNode == nil {
		return false
	}

	if value == currentNode.Value {
		return true
	}

	return false
}

func (tree *BST) Remove(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.
	return tree
}

func main() {
	bst := new(BST)
	bst.Value = 15
	bst = bst.Insert(2)
	fmt.Println(bst, bst.Left, bst.Right)
	bst = bst.Insert(3)
	fmt.Println(bst, bst.Left, bst.Right)
	bst = bst.Insert(4)
	fmt.Println(bst, bst.Left, bst.Right)

	response := bst.Contains(2)
	fmt.Println(response)
}
