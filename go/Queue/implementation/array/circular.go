package main

import "fmt"

type Queue struct {
	Elements []int
	Capacity int
	Size     int
	Front    int
}

func New(capacity int) *Queue {
	return &Queue{
		Elements: make([]int, capacity),
		Capacity: capacity,
		Size:     0,
	}
}

func (q *Queue) Enqueue(element int) {
	if q.IsFull() {
		return
	}

	rear := q.GetRear()
	fmt.Println("REAR:", rear)
	rear = (rear + 1) % q.Capacity
	q.Elements[rear] = element
	q.Size += 1
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	temp := q.Elements[q.Front]
	q.Elements[q.Front] = -1
	q.Front = (q.Front + 1) % q.Capacity
	q.Size -= 1

	return temp
}

func (q *Queue) GetFront() int {
	if q.IsEmpty() {
		return -1
	} else {
		return q.Front
	}
}

func (q *Queue) GetRear() int {
	if q.IsEmpty() {
		return -1
	} else {
		return (q.Front + q.Size - 1) % q.Capacity
	}
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) IsFull() bool {
	return q.Size == q.Capacity
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
