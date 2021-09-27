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

func BranchSums(node *Node) {
	count := 0
	BranchSumsHelper(node, count)
}
func BranchSumsHelper(node *Node, count int) {
	if node == nil { // When the node is nil
		return
	}

	count += node.Key
	if node.Left == nil && node.Right == nil {
		fmt.Printf("%d \t", count)
	}

	BranchSumsHelper(node.Left, count)
	BranchSumsHelper(node.Right, count)
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

	BranchSums(node)
}
