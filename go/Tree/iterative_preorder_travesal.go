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

func IterativePreOrderTraversal(tree *Tree) {
	if tree == nil {
		return
	}

	current := tree
	stack := NewStack()

	stack.Push(tree)
	for len(stack.Elements) > 0 {
		current = stack.Pop()
		fmt.Println(current.Value)

		if current.Right != nil {
			stack.Push(current.Right)
		}

		if current.Left != nil {
			stack.Push(current.Left)
		}
	}
}

func main() {
	node := &Tree{Value: 10}
	node.Left = &Tree{Value: 20}
	node.Right = &Tree{Value: 30}
	IterativePreOrderTraversal(node)
	fmt.Println()
}
