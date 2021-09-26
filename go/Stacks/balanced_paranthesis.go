package main

import "fmt"

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
	if s.IsEmpty() {
		return "" // Scenario when you do not allow
	}

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

// {} [] ()
// [{()}]
func BalancedParanthesis(input string) bool {

	// Initialize New stack
	stack := New()

	// Inputs to match closed brackets with Open ones
	matched := map[string]string{
		"}": "{",
		")": "(",
		"]": "[",
	}

	for _, char := range input {
		if string(char) == "{" || string(char) == "(" || string(char) == "[" {
			// Push the element stack
			stack.Push(string(char))
		} else if string(char) == "}" || string(char) == ")" || string(char) == "]" {
			elem := stack.Pop()
			if matched[string(char)] != elem {
				return false
			}
		}
	}

	fmt.Println(stack.Elements)
	return len(stack.Elements) == 0
}

func main() {
	s := "()[]{}"

	fmt.Println(BalancedParanthesis(s))
}
