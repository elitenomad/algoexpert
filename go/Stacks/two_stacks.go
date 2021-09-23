package main

import "fmt"

/*
	Two Stacks - LIFO Last in First Out
	Methods:
		Push
		Pop
		Peek

	Implementation:
		Array: Starting end defines First stack
					 Ending end defines Last stack
*/

type Stack struct {
	Elements []int
	topa     int
	topb     int
}

func New(size int) *Stack {
	return &Stack{
		Elements: make([]int, size),
		topa:     0,
		topb:     size - 1,
	}
}

func (s *Stack) PushA(element int) {
	if s.topa >= s.topb-1 {
		fmt.Println("Size reached its limit")
		return
	} else {
		s.Elements[s.topa] = element
		s.topa++
	}
}

func (s *Stack) PushB(element int) {
	if s.topb <= s.topa-1 {
		fmt.Println("Size reached its limit")
		return
	} else {
		s.Elements[s.topb] = element
		s.topb--
	}
}

func (s *Stack) PopA() int { // Ensure you handle failure usecase
	elemToBePopped := s.Elements[s.topa]
	s.Elements[s.topa] = -1
	s.topa--
	return elemToBePopped
}

func (s *Stack) PopB() int { // Ensure you handle failure usecase
	elemToBePopped := s.Elements[s.topb]
	s.Elements[s.topb] = -1
	s.topb++
	return elemToBePopped
}

func (s *Stack) PeekA() int {
	return s.Elements[s.topa-1]
}

func (s *Stack) PeekB() int {
	return s.Elements[s.topb+1]
}

func (s *Stack) IsEmptyA() bool {
	return s.topa <= 0
}

func (s *Stack) IsEmptyB() bool {
	return s.topb >= len(s.Elements)-1
}

func (s *Stack) SizeA() int {
	return s.topa
}
func (s *Stack) SizeB() int {
	return len(s.Elements) - s.topb - 1
}

func main() {
	s := New(10)
	fmt.Println(s.IsEmptyA())
	fmt.Println(s.IsEmptyB())
	fmt.Println()
	s.PushA(1)
	s.PushA(2)

	fmt.Println(s.SizeA())
	fmt.Println(s.SizeB())
	s.PushA(1)
	s.PushA(2)
	s.PushA(1)
	s.PushA(2)
	s.PushA(1)
	s.PushA(23)
	s.PushA(24)
	fmt.Println()
	fmt.Println(s.SizeA())
	fmt.Println(s.SizeB())

	s.PushB(3)
	s.PushB(4)
	s.PushB(5)
	fmt.Println(s.SizeA())
	fmt.Println(s.SizeB())
	fmt.Println()

	fmt.Println(s.PeekA())
	fmt.Println(s.PeekB())
	fmt.Println(s)
}
