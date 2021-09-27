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

func HeightOfSubtree(node *Node, count *int) int {
	if node == nil {
		return 0
	}

	*count += 1
	return HeightOfSubtree(node.Left, count) + HeightOfSubtree(node.Right, count)
}

func BalancedBinaryTree(node *Node) bool {
	if node == nil { // When the node is nil
		return true
	}

	count_l1 := 0
	count_l2 := 0
	HeightOfSubtree(node.Left, &count_l1)
	HeightOfSubtree(node.Right, &count_l2)

	fmt.Println(count_l1, count_l2)

	return count_l1 == count_l2
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

	fmt.Println(BalancedBinaryTree(node))
}
