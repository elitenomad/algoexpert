package main

import (
	"fmt"

	"github.com/elitenomad/godstructs/src/queue"
)

// type Node struct {
// 	Value string
// 	Next  *Node
// }
// type LinkedList struct {
// 	Head *Node
// 	Tail *Node
// }
// type Queue struct {
// 	Element *LinkedList
// }

// func NewQueue() *Queue {
// 	return &Queue{
// 		Element: &LinkedList{},
// 	}
// }

// func (q *Queue) Enqueue(element string) {
// 	node := &Node{Value: element}

// 	if q.Element.Head == nil {
// 		q.Element.Head = node
// 		q.Element.Tail = node
// 		return
// 	}

// 	q.Element.Tail.Next = node
// 	q.Element.Tail = node
// }

// func (q *Queue) Dequeue() *Node {
// 	nodeToBeRemoved := q.Element.Head
// 	q.Element.Head = q.Element.Head.Next

// 	nodeToBeRemoved.Next = nil
// 	return nodeToBeRemoved
// }

// func (q *Queue) IsEmpty() bool {
// 	return q.Element.Head == nil
// }

func GenerateNumbersII(n int) *queue.Queue {
	queue := queue.New()

	queue.Enqueue("5")
	queue.Enqueue("6")

	for i := 0; i < n; i++ {
		current := queue.Dequeue()
		fmt.Println(i, *current)

		queue.Enqueue(fmt.Sprintf("%v", *current) + "5")
		queue.Enqueue(fmt.Sprintf("%v", *current) + "6")
	}

	return queue
}

func main() {
	GenerateNumbersII(10)
}
