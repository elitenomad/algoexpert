package main

import "fmt"

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
	if s.IsEmpty() {
		return -1 // Scenario when you do not allow
	}

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

// {} [] ()
// [{()}]
// func BalancedParanthesis(input string) bool {

// 	// Initialize New stack
// 	stack := New()

// 	// Inputs to match closed brackets with Open ones
// 	matched := map[string]string{
// 		"}": "{",
// 		")": "(",
// 		"]": "[",
// 	}

// 	for _, char := range input {
// 		if string(char) == "{" || string(char) == "(" || string(char) == "[" {
// 			// Push the element stack
// 			stack.Push(string(char))
// 		} else if string(char) == "}" || string(char) == ")" || string(char) == "]" {
// 			elem := stack.Pop()
// 			if matched[string(char)] != elem {
// 				return false
// 			}
// 		}
// 	}

// 	fmt.Println(stack.Elements)
// 	return len(stack.Elements) == 0
// }

func BalancedParanthesisExtended(input string) bool {
	stack := New()
	stackTwo := New()

	for i, char := range input {
		if char == '(' {
			stack.Push(i)
			// fmt.Println(stack.Elements)
		} else if char == '*' {
			stackTwo.Push(i)
			// fmt.Println(stackTwo.Elements)
		} else if string(char) == ")" {

			elem := stack.Pop()

			if elem == -1 {
				if CompareItWithStarBeforeIndexes(stackTwo, i) {
					stackTwo.Pop()
				} else {
					return false
				}
			} else {
				// Do nothing as this condition is matched
			}
		}
	}

	for !stack.IsEmpty() && !stackTwo.IsEmpty() {
		elem := stack.Pop()

		if CompareItWithStarAfterIndexes(stackTwo, elem) {
			stackTwo.Pop()
		} else {
			stack.Push(elem)
		}
	}

	fmt.Println(stack.Elements)
	return len(stack.Elements) == 0
}

func CompareItWithStarBeforeIndexes(input *Stack, j int) bool {
	for i := 0; i < len(input.Elements); i++ {
		if input.Elements[i] < j {
			return true
		}
	}
	return false
}

func CompareItWithStarAfterIndexes(input *Stack, j int) bool {
	for i := 0; i < len(input.Elements); i++ {
		if input.Elements[i] > j {
			return true
		}
	}
	return false
}

func main() {
	s := "()"
	fmt.Println(BalancedParanthesisExtended(s))
	s = "*)"
	fmt.Println(BalancedParanthesisExtended(s))
	s = "(*)"
	fmt.Println(BalancedParanthesisExtended(s))
	s = "(*"
	fmt.Println(BalancedParanthesisExtended(s))
	s = "((*"
	fmt.Println(BalancedParanthesisExtended(s))
}
