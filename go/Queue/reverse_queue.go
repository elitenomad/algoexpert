package main

import "fmt"

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

func (q *Queue) GetFront() int {
	return q.Elements[0]
}

func (q *Queue) GetRear() int {
	return q.Elements[len(q.Elements)-1]
}

func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

func (q *Queue) ReverseQueue() {
	// This is typical implementation of an array reverse
	for i, j := 0, len(q.Elements)-1; i < len(q.Elements)/2; i, j = i+1, j-1 {
		fmt.Println(q.Elements[i], q.Elements[j])
		q.Elements[i], q.Elements[j] = q.Elements[j], q.Elements[i]
	}

	// Consider working with Queue methods
	/*
		for q.IsEmpty() {
			stack.push(q.deque())
		}

		for stack.IsEmpty() {
			queue.push(stack.pop())
		}

		return queue (Reversed with front and Rear reversed as well)
	*/
}

func main() {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.Size())
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())

	q.ReverseQueue()
	fmt.Println(q)
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
}
