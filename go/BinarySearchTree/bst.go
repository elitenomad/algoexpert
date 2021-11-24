package main

import (
	"fmt"
)

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func Init() *BST {
	return &BST{}
}

func Search(tree *BST, value int) bool {
	if tree == nil {
		return false
	}

	if tree.Value == value {
		return true
	} else if value < tree.Value {
		return Search(tree.Left, value)
	} else if value > tree.Value {
		return Search(tree.Right, value)
	}

	return false // Has to be false at this place.
}

func Insert(tree *BST, value int) *BST {
	if tree == nil {
		bst := Init()
		bst.Value = value
		return bst
	}

	if value < tree.Value {
		tree.Left = Insert(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = Insert(tree.Right, value)
	}

	return tree // Has to be false at this place.
}

func Delete(tree *BST, value int) *BST {
	if tree == nil {
		return nil
	}

	if value < tree.Value {
		tree.Left = Insert(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = Insert(tree.Right, value)
	} else {
		if tree.Left == nil {
			return tree.Right
		} else if tree.Right == nil {
			return tree.Left
		} else {
			successor := getSuccessor(tree)
			tree.Value = successor.Value
			tree.Right = Delete(tree.Right, successor.Value)
		}
	}

	return tree
}

func getSuccessor(root *BST) *BST {
	curr := root.Right
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func Floor(root *BST, value int) int { // Closest small or equal value
	curr := root

	if curr.Value == value {
		return curr.Value
	}

	for curr.Left != nil || curr.Right != nil {
		if value < curr.Value {
			curr = curr.Left
		} else if value > curr.Value {
			curr = curr.Right
		} else {
			return value
		}
	}

	return curr.Value
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func Ceil(root *BST, value int) int { // Closest large or equal value
	curr := root
	var result *BST

	for curr != nil {
		if value < curr.Value {
			result = curr
			curr = curr.Left
		} else if value > curr.Value {
			curr = curr.Right
		} else {
			return value
		}
	}

	return result.Value
}

func main() {
	bst := &BST{Value: 10}
	bst.Left = &BST{Value: 5}
	bst.Right = &BST{Value: 15}
	bst.Right.Left = &BST{Value: 12}
	bst.Right.Right = &BST{Value: 30}
	// Insert(bst, 10)
	// Insert(bst, 5)
	// Insert(bst, 15)
	// Insert(bst, 12)
	// Insert(bst, 18)
	// fmt.Println(Search(bst, 15))
	// fmt.Println(Search(bst, 100))
	// fmt.Println(Search(bst, 5))
	// bst = Delete(bst, 5)
	// fmt.Println(bst.Left)
	// fmt.Println(Search(bst, 5))
	fmt.Println(Floor(bst, 14))
	fmt.Println(Floor(bst, 30))
	fmt.Println(Ceil(bst, 14))
	fmt.Println(Ceil(bst, 30))
}
