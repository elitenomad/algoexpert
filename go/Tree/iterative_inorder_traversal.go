package main

import "fmt"

type Stack struct {
	Elements []*Tree
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(tree *Tree) {
	s.Elements = append(s.Elements, tree)
}

func (s *Stack) Pop() *Tree {
	poppedElement := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return poppedElement
}

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func IterativeInOrderTraversal(tree *Tree) {
	current := tree
	stack := NewStack()

	for current != nil || len(stack.Elements) > 0 {
		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		current = stack.Pop()
		fmt.Println(current.Value)
		current = current.Right
	}
}

func main() {
	node := &Tree{Value: 10}
	node.Left = &Tree{Value: 20}
	node.Right = &Tree{Value: 30}
	IterativeInOrderTraversal(node)
	fmt.Println()
}
