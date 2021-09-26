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

func FindMax(node *Node) int {
	max := 0
	return InOrderTraversalHelper(node, &max)
}

func InOrderTraversalHelper(node *Node, max *int) int {
	if node == nil {
		return -1
	}

	if node.Key > *max {
		*max = node.Key
	}

	InOrderTraversalHelper(node.Left, max)
	InOrderTraversalHelper(node.Right, max)

	return *max
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
	fmt.Println(FindMax(node))
}
