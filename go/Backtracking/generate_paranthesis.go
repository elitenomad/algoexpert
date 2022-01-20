package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	Elements []string
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(element string) {
	s.Elements = append(s.Elements, element)
}

func (s *Stack) Pop() string {
	elemToBePopped := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return elemToBePopped
}

func (s *Stack) Peek() string {
	return s.Elements[len(s.Elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.Elements) <= 0
}

func (s *Stack) Size() int {
	return len(s.Elements)
}

func generateParanthesis(n int) []string {
	// Only add open paranthesis if open < n
	// Only add closing paranthesis if closed < open
	// Valid if open == closed == n
	result := []string{}
	stack := New()

	generateParanthesisHelper(n, 0, 0, stack, &result)
	return result
}

func generateParanthesisHelper(n int, open, closed int, stack *Stack, result *[]string) {
	if open == n && closed == n {
		*result = append(*result, strings.Join(stack.Elements, ""))
		return
	}

	if open < n {
		stack.Push("(")
		generateParanthesisHelper(n, open+1, closed, stack, result)
		stack.Pop()
	}

	if closed < open {
		stack.Push(")")
		generateParanthesisHelper(n, open, closed+1, stack, result)
		stack.Pop()
	}

}

func main() {
	fmt.Println(generateParanthesis(1))
}
