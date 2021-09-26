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

func CountNoOfNodes(node *Node) int {
	count := 0 // Initialize a counter

	// Use InOrder Traversal ROOT LEFT RIGHT DFS
	// Send a reference to a count variable
	// Update the count whenever it enter its recursive
	return InOrderTraversalHelper(node, &count)
}

func InOrderTraversalHelper(node *Node, count *int) int {
	if node == nil {
		return 0
	}

	*count += 1
	InOrderTraversalHelper(node.Left, count)
	InOrderTraversalHelper(node.Right, count)

	return *count
}

// FIND THE SIZE OF BINARY TREE - Below code counts
// The number of nodes using Level Ordering Traversal
func UseLevelOrderingTraversal(node *Node) int {
	// queue := NewQueue()
	// count := 0

	// queue.Enqueue(node)

	// for !queue.IsEmpty() {
	// 	size := len(queue.Elements)
	// 	count += size
	// 	fmt.Println(count)
	// 	for i := 0; i < size; i++ {
	// 		qd := queue.Dequeue()

	// 		if qd.Left != nil {
	// 			queue.Enqueue(qd.Left)
	// 		}
	// 		if qd.Right != nil {
	// 			queue.Enqueue(qd.Right)
	// 		}
	// 	}
	// }

	return -1 // Return count
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
	fmt.Println(CountNoOfNodes(node))
}
