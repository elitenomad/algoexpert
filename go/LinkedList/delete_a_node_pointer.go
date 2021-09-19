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

// Given a Pointer and no information about Head of the Linked List
// You need to delete the pointer from the Linked List

// Node to be deleted will never be the Last node - Assumption we will go with
func (l *LinkedList) DeleteANodeGivenAPointer(pointer *Node) {
	current := pointer

	current.Value = current.Next.Value
	current.Next = current.Next.Next

	// Above two lines will suffice :) ha ha
	// Below logic is just to get a mental map of how to achieve
	// what we achieve above with 2 steps. (Below goes through the loop until the end of LinkedList)
	// for current.Next != nil {
	// 	previous = current
	// 	current = current.Next

	// 	if current != nil {
	// 		previous.Value = current.Value
	// 	}
	// 	previous.Next = current
	// }

	// previous.Next = nil
	// current = nil
}

// variation Question
// Length of the Loop
// First node of the Loop

func main() {
	ll := NewLinkedList()

	ll.Append(10)
	ll.Append(15)
	ll.Append(12)
	ll.Append(20)

	// ll.CreateCycle(15, 20)
	node := ll.ReturnANode(15) // Fetch the value of Node using regular process
	ll.DeleteANodeGivenAPointer(node)
	ll.List()
}
