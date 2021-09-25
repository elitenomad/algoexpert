package main

import "fmt"

type Node struct {
	Value string
	Next  *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
}
type Queue struct {
	Element *LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		Element: &LinkedList{},
	}
}

func (q *Queue) Enqueue(element string) {
	node := &Node{Value: element}

	if q.Element.Head == nil {
		q.Element.Head = node
		q.Element.Tail = node
		return
	}

	q.Element.Tail.Next = node
	q.Element.Tail = node
}

func (q *Queue) Dequeue() *Node {
	nodeToBeRemoved := q.Element.Head
	q.Element.Head = q.Element.Head.Next

	nodeToBeRemoved.Next = nil
	return nodeToBeRemoved
}

func (q *Queue) IsEmpty() bool {
	return q.Element.Head == nil
}

func GenerateNumbers(n int) Queue {
	queue := NewQueue()

	queue.Enqueue("5")
	queue.Enqueue("6")

	for i := 0; i < n; i++ {
		current := queue.Dequeue()
		fmt.Println(i, current.Value)
		queue.Enqueue(current.Value + "5")
		queue.Enqueue(current.Value + "6")
	}

	return *queue
}

func main() {
	GenerateNumbers(10)
}
