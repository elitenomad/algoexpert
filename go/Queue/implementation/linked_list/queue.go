package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Append(value int) {
	node := &Node{Value: value}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = node
}

type Queue struct {
	List *LinkedList
}

func New(capacity int) *Queue {
	return &Queue{List: NewLinkedList()}
}

func (q *Queue) Enqueue(element int) {
	q.List.Append(element)
}

func (q *Queue) Dequeue() *Node {
	if q.List.Head == nil {
		return nil
	}

	temp := q.List.Head
	q.List.Head = temp.Next

	return temp
}

func (q *Queue) GetFront() int {
	if q.List.Head == nil {
		return -1
	} else {
		return q.List.Head.Value
	}
}

func (q *Queue) GetRear() int {
	if q.List.Tail == nil {
		return -1
	} else {
		return q.List.Tail.Value
	}
}

func main() {
	q := New(10)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
	fmt.Println(q)
	q.Dequeue()
	q.Dequeue()

	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
}
