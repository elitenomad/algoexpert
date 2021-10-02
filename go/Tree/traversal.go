package main

import "fmt"

/*
	Tree Traversal
		- Breadth First (Level Order)
		- Depth First
			- In Order
				- LEFT ROOT RIGHT
			- Pre Order
				- ROOT LEFT RIGHT
			- Post Order
				-
*/

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func New() *Node {
	return &Node{}
}

// Depth First Algorithms
// InOrder Traversal - LEFT ROOT RIGHT
func InOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	InOrderTraversal(node.Left)
	fmt.Println(node.Key)
	InOrderTraversal(node.Right)
}

// Pre order Traversal - Root LEFT RIGHT
func PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Key)
	InOrderTraversal(node.Left)
	InOrderTraversal(node.Right)
}

// Post Order Traversal - LEFT RIGHT ROOT
func PostOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	InOrderTraversal(node.Left)
	InOrderTraversal(node.Right)
	fmt.Println(node.Key)
}

func main() {
	node := &Node{Key: 10}
	node.Left = &Node{Key: 20}
	node.Right = &Node{Key: 30}
	InOrderTraversal(node)
	fmt.Println()
	PreOrderTraversal(node)
	fmt.Println()
	PostOrderTraversal(node)
}
