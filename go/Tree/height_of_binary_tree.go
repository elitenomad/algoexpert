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

func UsePreOrderTraversalToFindHeightOfBT(node *Node) int {
	return PreOrderTraversalHelper(node, 0)
}

func PreOrderTraversalHelper(node *Node, height int) int {
	if node == nil {
		return height
	}

	height = height + 1
	h1 := PreOrderTraversalHelper(node.Left, height)
	h2 := PreOrderTraversalHelper(node.Right, height)

	return max(h1, h2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	node := New()
	node.Key = 10
	node.Left = New()
	node.Left.Key = 5
	node.Right = New()
	node.Right.Key = 20
	node.Right.Left = New()
	node.Right.Right = New()
	node.Right.Left.Key = 15
	node.Right.Right.Key = 40
	fmt.Println(UsePreOrderTraversalToFindHeightOfBT(node))
}
