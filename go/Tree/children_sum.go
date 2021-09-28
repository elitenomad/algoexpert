package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func New() *Node {
	return &Node{}
}

// Using InOrder Traversal to get Children Sum
// Tree
func ChildrenSum(node *Node) bool {
	if node == nil { // When the node is nil
		return true
	}

	if node.Left == nil && node.Right == nil { // When it has only one node
		return true
	}

	// Main use case for the root
	// Initialise variable sum to 0
	// Reduce the sum value with both left and right tree Nodes given they are not nil
	// Equate it with root sum and recursiviely call Left and Right Nodes
	sum := 0
	if node.Left != nil {
		sum += node.Left.Key
	}

	if node.Right != nil {
		sum += node.Right.Key
	}

	return (node.Key == sum) && ChildrenSum(node.Left) && ChildrenSum(node.Right)
}

func main() {
	node := New()
	node.Key = 1
	node.Left = New()
	node.Left.Key = 2
	node.Right = New()
	node.Right.Key = 3
	node.Left.Left = New()
	node.Left.Right = New()
	node.Left.Left.Key = 4
	node.Left.Right.Key = 24
	node.Right.Left = New()
	node.Right.Right = New()
	node.Right.Left.Key = 6
	node.Right.Right.Key = 7

	fmt.Println(ChildrenSum(node))
}
