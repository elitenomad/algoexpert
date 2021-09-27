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

// Using LevelOrder Traversal Print the Left View of the
// Tree
func PrintLeftViewOfATree(node *Node) {
	queue := NewQueue()

	queue.Enqueue(node)

	for !queue.IsEmpty() {
		size := len(queue.Elements)

		// Always collect the first of the Queued elements at every Leve
		// And display it
		fmt.Printf("%d \t", queue.Elements[0].Key)

		for i := 0; i < size; i++ {
			dequed := queue.Dequeue()

			if dequed.Left != nil {
				queue.Enqueue(dequed.Left)
			}

			if dequed.Right != nil {
				queue.Enqueue(dequed.Right)
			}
		}
	}
}

// Recursive way using PreOrder Traversal
func PrintLeftViewOfATreeHelper(node *Node, maxLevel *int, level int) {
	if node == nil {
		return
	}

	if *maxLevel < level {
		fmt.Printf("%d \t", node.Key)
		*maxLevel = level
	}

	PrintLeftViewOfATreeHelper(node.Left, maxLevel, level+1)
	PrintLeftViewOfATreeHelper(node.Right, maxLevel, level+1)
}

func PrintLeftViewOfATreeII(node *Node) {
	maxLevel := 0
	PrintLeftViewOfATreeHelper(node, &maxLevel, 1)
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

	PrintLeftViewOfATree(node)
	fmt.Println()
	PrintLeftViewOfATreeII(node)
}
