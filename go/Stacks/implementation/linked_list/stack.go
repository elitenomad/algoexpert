package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

type Stack struct {
	element *LinkedList
}

func NewStack() *Stack {
	return &Stack{
		element: &LinkedList{},
	}
}

func (l *LinkedList) Prepend(value int) {
	node := &Node{Value: value, Next: l.Head}
	l.Head = node
	l.Size += 1
}

func (l *LinkedList) DeleteTail() *Node {
	head := l.Head
	l.Head = l.Head.Next
	l.Size -= 1

	return head
}

// Stack Methods

func (s *Stack) Push(element int) {
	s.element.Prepend(element)
}

func (s *Stack) Pop() int {
	return s.element.DeleteTail().Value
}

func (s *Stack) Peek() int {
	return s.element.Head.Value
}

func (s *Stack) Size() int {
	return s.element.Size
}

func (s *Stack) IsEmpty() bool {
	return s.element.Size <= 0
}

func main() {
	s := NewStack()
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Size())
	s.Pop()
	fmt.Println(s.Size())
	fmt.Println(s.IsEmpty())
	s.Pop()
	s.Pop()
	fmt.Println(s.Size())
	fmt.Println(s.IsEmpty())
}
