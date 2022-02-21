package main

import "fmt"

/*
Design a DS where these operation are performed efficiently
	Search
	Insert
	Delete
	K-th Smallest
*/

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

// Using InOrder traversal
func KthSmallest(tree *BST, k int) int {
	count := 0
	result := 0
	KthSmallestHelper(tree, k, &count, &result)
	return result
}

func KthSmallestHelper(tree *BST, k int, count *int, result *int) {
	if tree == nil {
		return
	}

	KthSmallestHelper(tree.Left, k, count, result)
	*count += 1
	fmt.Println(count, k, tree.Value)
	if *count == k {
		*result = tree.Value
		return
	}
	KthSmallestHelper(tree.Right, k, count, result)
}

func main() {
	root := &BST{}
	root.Value = 15
	root.Left = &BST{Value: 5}
	root.Left.Left = &BST{Value: 3}
	root.Right = &BST{Value: 20}
	root.Right.Left = &BST{Value: 18}
	root.Right.Left.Left = &BST{Value: 16}
	root.Right.Right = &BST{Value: 80}
	k := 3
	fmt.Println(KthSmallest(root, k))
}
