package main

import "fmt"

/*
	Queue DS - FIFO First IN First Out
	Methods:
		Enqueue
		Dequeue

	Implementation:
		Array
*/

type Queue struct {
	Elements []int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(element int) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() int { // O(N)
	removedElement := q.Elements[0]
	q.Elements = q.Elements[1:]
	return removedElement
}

func (q *Queue) Front() int {
	return q.Elements[0]
}

func (q *Queue) Rear() int {
	return q.Elements[len(q.Elements)-1]
}

func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

func main() {
	q := New()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Size())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Size())
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Size())
	fmt.Println(q.Front())
	fmt.Println(q.Rear())
}
