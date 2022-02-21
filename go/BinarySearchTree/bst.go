package main

import (
	"fmt"
	"math"
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

func SearchIterative(tree *BST, value int) bool { // O(h) and O(1)
	for tree != nil {
		if tree.Value == value {
			return true
		} else if value < tree.Value {
			tree = tree.Left
		} else if value > tree.Value {
			tree = tree.Right
		}
	}

	return false
}

func Search(tree *BST, value int) bool { // O(h) and O(h) auxillary stack space because of recursive
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

	return false
}

func InsertIterative(tree *BST, value int) *BST {
	if tree == nil {
		bst := Init()
		bst.Value = value
		return bst
	}

	current := tree
	for current != nil {
		if value == current.Value {
			return tree
		} else if value < current.Value {
			if current.Left == nil {
				current.Left = &BST{Value: value}
			} else {
				current = current.Left
			}
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = &BST{Value: value}
			} else {
				current = current.Right
			}
		}
	}

	return tree
}

func InsertIterativeII(tree *BST, value int) *BST {
	bst := Init()
	bst.Value = value

	current := tree
	var parent *BST

	for current != nil {
		parent = current
		if value == current.Value {
			return tree
		} else if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		}
	}

	if value < parent.Value {
		parent.Left = bst
	} else if value > parent.Value {
		parent.Right = bst
	}

	return tree
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
		tree.Left = Delete(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = Delete(tree.Right, value)
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
	var res *BST

	for curr != nil {
		if value < curr.Value {
			curr = curr.Left
		} else if value > curr.Value {
			res = curr
			curr = curr.Right
		} else {
			return value
		}
	}

	return res.Value
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

func CheckBST(tree *BST) bool {
	return CheckBSTHelper(tree, math.MinInt32, math.MaxInt32)
}

/*
				80 (-INF to INF)
		(-INF to 80)
		70 				90(80 to INF)
(-INF to 70)(70 to 80)(80 to 90)
	60		75		85		100 (90 to INF)
*/
func CheckBSTHelper(tree *BST, min, max int) bool {
	if tree == nil {
		return true
	}

	if tree.Value < min || tree.Value >= max {
		return false
	}

	leftValid := CheckBSTHelper(tree.Left, min, tree.Value)
	rightValid := CheckBSTHelper(tree.Right, tree.Value, max)

	return leftValid && rightValid
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
