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

	return abs(count_l1-count_l2) <= 1
}

func BalancedBinaryTreeII(root *Node) int {
	if root == nil {
		return 0
	}

	lh := BalancedBinaryTreeII(root.Left)
	if lh == -1 {
		return -1
	}
	rh := BalancedBinaryTreeII(root.Right)
	if rh == -1 {
		return -1
	}

	if abs(lh-rh) > 1 {
		return -1
	} else {
		return max(lh, rh) + 1
	}
}

func abs(s int) int {
	if s < 0 {
		return -s
	}

	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	node := New()
	node.Key = 1
	node.Left = New()
	node.Left.Key = 2
	// node.Right = New()
	// node.Right.Key = 3
	node.Left.Left = New()
	node.Left.Right = New()
	node.Left.Left.Key = 4
	node.Left.Right.Key = 24
	// node.Right.Left = New()
	// node.Right.Right = New()
	// node.Right.Left.Key = 6
	// node.Right.Right.Key = 7

	fmt.Println(BalancedBinaryTree(node))

	if BalancedBinaryTreeII(node) > 0 {
		fmt.Println("Its Balanced")
	} else {
		fmt.Println("Its not Balanced")
	}
}
