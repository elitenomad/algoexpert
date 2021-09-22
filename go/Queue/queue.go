package main

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

func (q *Queue) Dequeue() int {
	removedElement := q.Elements[0]
	q.Elements = q.Elements[1:]
	return removedElement
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

func main() {

}
