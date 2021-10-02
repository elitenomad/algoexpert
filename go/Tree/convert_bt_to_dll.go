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

func ConvertBinaryTreetoDoublyLinkedList(node *Node) *Node {
	var prev *Node
	return ConvertBinaryTreetoDoublyLinkedListHelper(node, prev)
}

func ConvertBinaryTreetoDoublyLinkedListHelper(node *Node, prev *Node) *Node {
	if node == nil {
		return node
	}

	head := ConvertBinaryTreetoDoublyLinkedListHelper(node.Left, prev)
	if prev == nil {
		head = node
	} else {
		node.Left = prev
		prev.Right = node
	}

	prev = node
	ConvertBinaryTreetoDoublyLinkedListHelper(node.Right, prev)
	return head
}

func main() {
	node := New()
	node.Key = 10
	node.Left = New()
	node.Left.Key = 5
	node.Right = New()
	node.Right.Key = 20
	// node.Left.Left = New()
	// node.Left.Right = New()
	// node.Left.Left.Key = 4
	// node.Left.Right.Key = 24
	node.Right.Left = New()
	node.Right.Right = New()
	node.Right.Left.Key = 30
	node.Right.Right.Key = 35

	result := ConvertBinaryTreetoDoublyLinkedList(node)

	current := result
	fmt.Println(current.Key)
	for current != nil {
		fmt.Println("[", current.Key, "]")
		current = current.Left
	}
}
