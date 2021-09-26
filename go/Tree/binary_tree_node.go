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

func main() {
	node := New()
	node.Key = 10
	node.Left = New()
	node.Left.Key = 5
	node.Right = New()
	node.Right.Key = 20
	fmt.Println(node)
}
