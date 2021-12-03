package main

import "fmt"

/*
	Stack implementation
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

func PreviousSmallArray(input []int) []int {
	result := make([]int, len(input))
	s := New()

	s.Push(0)
	result[0] = -1
	for i := 0; i < len(input); i++ {
		for !s.IsEmpty() && input[i] <= input[s.Peek()] {
			s.Pop()
		}
		var res int
		if s.IsEmpty() {
			res = -1
		} else {
			res = s.Pop()
		}
		s.Push(i)
		result[i] = res
	}

	return result
}

func NextSmallArray(input []int) []int {
	result := make([]int, len(input))
	s := New()

	s.Push(len(input) - 1)
	result[len(input)-1] = -1
	for i := len(input) - 1; i >= 0; i-- {
		for !s.IsEmpty() && input[i] <= input[s.Peek()] {
			s.Pop()
		}
		var res int
		if s.IsEmpty() {
			res = -1
		} else {
			res = s.Pop()
		}
		s.Push(i)
		result[i] = res
	}

	return result
}

func LargestRectangularAreaUsingStacks(input []int) int {
	ps := PreviousSmallArray(input)
	ns := NextSmallArray(input)
	fmt.Println(ps, ns)

	res := 0
	for i := 0; i < len(input); i++ {
		curr := input[i]
		if ps[i] >= 0 {
			curr += (i - ps[i] - 1) * input[i]
		}

		if ns[i] >= 0 {
			curr += (ns[i] - i - 1) * input[i]
		}

		res = max(res, curr)
	}

	return res
}

// This is O(N^2)
func LargestRectangularArea(input []int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		current := input[i]
		for j := i - 1; j >= 0; j-- {
			if input[j] >= input[i] {
				current += input[i]
			} else {
				break
			}
		}

		for j := i + 1; j < len(input); j++ {
			if input[j] >= input[i] {
				current += input[i]
			} else {
				break
			}
		}

		result = max(result, current)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(LargestRectangularArea([]int{6, 2, 5, 4, 1, 5, 6}))
	fmt.Println(LargestRectangularAreaUsingStacks([]int{6, 2, 5, 4, 1, 5, 6}))
}
