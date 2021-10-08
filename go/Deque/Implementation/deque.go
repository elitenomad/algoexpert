package main

import "fmt"

/*
	GetFront()
	GetRear()
	IsFull()
	IsEmpty()
	IsSize()

	InsertFront
	DeleteFront
	InsertRear
	DeleteRear
*/

type DeQueue struct {
	Elements []int
	Front    int
	Rear     int
}

func New() *DeQueue {
	return &DeQueue{}
}

func (q *DeQueue) InsertRear(element int) {
	q.Rear = element
	if q.Size() <= 0 {
		q.Front = element
	}
	q.Elements = append(q.Elements, element)
}

func (q *DeQueue) DeleteRear() int { // O(N)

	if q.Size() <= 0 {
		return -1
	}

	if q.Size() == 1 {
		deltedElement := q.Elements[0]
		q.Front = -1
		q.Rear = -1
		q.Elements = []int{}
		return deltedElement
	}

	removedElement := q.Elements[q.Size()-1]
	q.Elements = q.Elements[:q.Size()-1]
	q.Rear = q.Elements[q.Size()-1]
	return removedElement
}

func (q *DeQueue) InsertFront(element int) {
	q.Front = element
	if q.Size() <= 0 {
		q.Rear = element
	}
	q.Elements = append([]int{element}, q.Elements...)
}

func (q *DeQueue) DeleteFront() int {
	if q.Size() <= 0 {
		return -1
	}

	if q.Size() == 1 {
		deltedElement := q.Elements[0]
		q.Front = -1
		q.Rear = -1
		q.Elements = []int{}
		return deltedElement
	}

	removedElement := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Front = q.Elements[0]

	return removedElement
}

func (q *DeQueue) GetFront() int {
	return q.Elements[0]
}

func (q *DeQueue) GetRear() int {
	return q.Elements[len(q.Elements)-1]
}

func (q *DeQueue) Size() int {
	return len(q.Elements)
}

func (q *DeQueue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

func main() {
	q := New()
	q.InsertFront(10)
	q.InsertRear(20)
	q.DeleteFront()
	fmt.Println(q.Size(), q.Front, q.Rear)
}
