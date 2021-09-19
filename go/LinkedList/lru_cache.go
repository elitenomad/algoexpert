package main

import "fmt"

type Node struct {
	Value   int
	Next    *Node
	Visited bool
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		Next:  nil,
	}
}

func (l *LinkedList) Append(value int) {
	node := NewNode(value)

	// Verify if there is a Head
	if l.Head == nil {
		l.Head = node

		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
}

func (l *LinkedList) ReturnANode(val int) *Node {
	current := l.Head

	for current != nil {
		if current.Value == val {
			return current
		}

		current = current.Next
	}

	return nil
}

func (l *LinkedList) List() {
	current := l.Head

	fmt.Println()
	for current != nil {
		fmt.Printf("Value : %d \t", current.Value)
		current = current.Next
	}
}

func main() {
	ll := NewLinkedList()

	ll.Append(5)
	ll.Append(8)
	ll.Append(7)
	ll.Append(10)
	ll.Append(12)
	ll.Append(15)

	ll.List()
}
