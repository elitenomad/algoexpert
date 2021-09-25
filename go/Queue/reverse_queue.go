package main

import "fmt"

// To reverse a Queue one of the appraoches is to use stack
// Building a stack structure to ensure we solve that problem
type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

type Stack struct {
	Element *LinkedList
}

func NewStack() *Stack {
	return &Stack{
		Element: &LinkedList{},
	}
}

func (s *Stack) Push(element int) {
	node := &Node{Value: element}

	if s.Element.Head == nil {
		s.Element.Head = node
		return
	}

	node.Next = s.Element.Head
	s.Element.Head = node
}

func (s *Stack) Pop() *Node {
	if s.Element.Head == nil {
		return nil
	} // Handles a use case when head is nil

	elementToBePopped := s.Element.Head
	s.Element.Head = s.Element.Head.Next
	return elementToBePopped // We can return a node here as well
}

func (s *Stack) IsEmpty() bool {
	return s.Element.Head == nil
}

// These method implemenations are good enough to solve the reverse queue problem

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

// This is typical implementation of an array reverse
func (q *Queue) ReverseQueueBasic() {
	for i, j := 0, len(q.Elements)-1; i < len(q.Elements)/2; i, j = i+1, j-1 {
		fmt.Println(q.Elements[i], q.Elements[j])
		q.Elements[i], q.Elements[j] = q.Elements[j], q.Elements[i]
	}
}

func (q *Queue) Reverse() {
	stack := NewStack()

	for !q.IsEmpty() {
		dequeuedElement := q.Dequeue()
		stack.Push(dequeuedElement)
	}

	for !stack.IsEmpty() {
		poppedElement := stack.Pop()
		fmt.Println(poppedElement.Value)
		if poppedElement != nil {
			q.Enqueue(poppedElement.Value)
		}
	}
}

func main() {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
	fmt.Println()
	q.Reverse()
	fmt.Println(q)
	fmt.Println(q.GetFront())
	fmt.Println(q.GetRear())
}
