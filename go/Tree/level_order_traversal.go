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

// Goal is to traverse the tree
// Level by Level
// Breadth First search || Level order traversal
func LeveOrderTraversal(node *Node) {
	queue := NewQueue()

	queue.Enqueue(node)

	for !queue.IsEmpty() {
		elem := queue.Dequeue()
		fmt.Println(elem.Key)

		if elem.Left != nil {
			queue.Enqueue(elem.Left)
		}

		if elem.Right != nil {
			queue.Enqueue(elem.Right)
		}
	}
}

func LeveOrderTraversalLineByLine(node *Node) {
	if node == nil {
		return
	}
	queue := NewQueue()
	dummyNode := &Node{Key: -1}

	queue.Enqueue(node)
	queue.Enqueue(dummyNode)

	fmt.Println(queue)
	for len(queue.Elements) > 1 {
		elem := queue.Dequeue()

		if elem.Key == -1 {
			fmt.Println()
			queue.Enqueue(dummyNode)
			continue
		}

		fmt.Printf("%d\t", elem.Key)

		if elem.Left != nil {
			queue.Enqueue(elem.Left)
		}

		if elem.Right != nil {
			queue.Enqueue(elem.Right)
		}
	}
}

func LeveOrderTraversalReturnArray(node *Node) [][]int {
	if node == nil {
		return [][]int{} // If root is nil return empty array
	}

	// Initialize DS required
	queue := NewQueue()
	result := [][]int{}

	// Add the root node to the queue
	queue.Enqueue(node)
	result = append(result, []int{node.Key})

	for !queue.IsEmpty() {
		size := len(queue.Elements)
		temp := []int{}

		for i := 0; i < size; i++ {
			elem := queue.Dequeue()
			if elem.Left != nil {
				queue.Enqueue(elem.Left)
				temp = append(temp, elem.Left.Key)
			}

			if elem.Right != nil {
				queue.Enqueue(elem.Right)
				temp = append(temp, elem.Right.Key)
			}
		}

		if len(temp) > 0 {
			result = append(result, temp)
		}
	}

	return result
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

	// LeveOrderTraversal(node)
	fmt.Println(LeveOrderTraversalReturnArray(node))
	// LeveOrderTraversalLineByLine(node)
}
