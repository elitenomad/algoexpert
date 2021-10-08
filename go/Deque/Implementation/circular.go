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
	Size     int
	Cap      int
	Front    int
}

func New(capacity int) *DeQueue {
	return &DeQueue{
		Elements: make([]int, capacity),
		Size:     0,
		Cap:      capacity,
		Front:    0,
	}
}

// Rear = (Rear + 1) % Cap
func (q *DeQueue) InsertRear(element int) {
	if q.IsFull() {
		return
	}

	newRear := (q.Front + q.Size) % q.Cap
	q.Elements[newRear] = element
	q.Size += 1
}

// Rear = (Rear - 1 + Cap) % Cap
func (q *DeQueue) DeleteRear() int { // O(N)
	if q.IsEmpty() {
		return -1
	}

	removedElement := q.Elements[q.GetRear()]
	q.Elements[q.GetRear()] = -1 // Ensure the value is set to -1 for differentiating
	q.Size -= 1

	return removedElement
}

// Front = (Front - 1 + Cap) % Cap
func (q *DeQueue) InsertFront(element int) {
	if q.IsFull() {
		return
	}

	q.Front = (q.Front - 1 + q.Cap) % q.Cap
	q.Elements[q.Front] = element
	q.Size++
}

// Front = (Front + 1) % Cap
func (q *DeQueue) DeleteFront() int {
	if q.IsEmpty() {
		return -1
	}

	removedElement := q.Elements[q.Front]
	q.Front = (q.Front + 1) % q.Cap
	q.Size -= 1

	return removedElement
}

func (q *DeQueue) GetFront() int {
	if q.IsEmpty() {
		return -1
	}

	return q.Front
}

// Rear = (Front + Size - 1) % Cap
func (q *DeQueue) GetRear() int {
	if q.IsEmpty() {
		return -1
	}

	return (q.Front + q.Size - 1) % q.Cap
}

func (q *DeQueue) IsEmpty() bool {
	return q.Size == 0
}

func (q *DeQueue) IsFull() bool {
	return q.Size == q.Cap
}

func main() {
	q := New(10)
	q.InsertFront(10)
	q.InsertRear(20)
	q.DeleteFront()
	// q.DeleteRear()
	fmt.Println(q.Size, q.GetFront(), q.GetRear())
}
