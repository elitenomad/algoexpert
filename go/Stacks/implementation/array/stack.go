package main

import "fmt"

/*
	Stack - LIFO Last in First Out
	Methods:
		Push
		Pop
		Peek

	Implementation:
		Array
*/

type Stack struct {
	Elements []int
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element int) {
	s.Elements = append(s.Elements, element)
}

func (s *Stack) Pop() int {
	elemToBePopped := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return elemToBePopped
}

func (s *Stack) Peek() int {
	return s.Elements[len(s.Elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) <= 0
}

func (s *Stack) Size() int {
	return len(s.Elements)
}

func main() {
	s := New()
	fmt.Println(s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s)
}
