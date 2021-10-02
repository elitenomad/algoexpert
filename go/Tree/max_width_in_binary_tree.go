package main

import "fmt"

//Implement a Queue
type Queue struct {
	Elements []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(node *Node) {
	q.Elements = append(q.Elements, node)
}

func (q *Queue) Dequeue() *Node {
	dequeuedNode := q.Elements[0]
	q.Elements = q.Elements[1:]
	return dequeuedNode
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

// Complete

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func New() *Node {
	return &Node{}
}

// Using the Level Traversal approach
func MaxWidthOfBT(root *Node) int {
	if root == nil {
		return 0
	}

	queue := NewQueue()
	queue.Enqueue(root)
	max := 0
	for !queue.IsEmpty() {
		size := len(queue.Elements)
		if max < size {
			max = size
		}

		for i := 0; i < size; i++ {
			deque := queue.Dequeue()
			if deque.Left != nil {
				queue.Enqueue(deque.Left)
			}

			if deque.Right != nil {
				queue.Enqueue(deque.Right)
			}
		}
	}

	return max
}

func main() {
	node := New()
	node.Key = 1
	node.Left = New()
	node.Left.Key = 2
	node.Right = New()
	node.Right.Key = 3
	// node.Left.Left = New()
	// node.Left.Right = New()
	// node.Left.Left.Key = 4
	// node.Left.Right.Key = 24
	// node.Right.Left = New()
	// node.Right.Right = New()
	// node.Right.Left.Key = 6
	// node.Right.Right.Key = 7
	fmt.Println(MaxWidthOfBT(node))
}
