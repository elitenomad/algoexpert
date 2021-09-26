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

func KthNodeFromRoot(node *Node, k int) { // Print the nodes when K
	InOrderTraversalHelper(node, k+1)
	// We send k + 1 as its distance from root to the nodes which should be displayed
}

func InOrderTraversalHelper(node *Node, k int) {
	if node == nil {
		return
	}

	k -= 1
	if k == 0 {
		fmt.Println(node.Key)
		return
	}
	InOrderTraversalHelper(node.Left, k)
	InOrderTraversalHelper(node.Right, k)

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
	node.Left.Right.Key = 5
	node.Right.Left = New()
	node.Right.Right = New()
	node.Right.Left.Key = 6
	node.Right.Right.Key = 7

	KthNodeFromRoot(node, 2)
}
