package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
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

func (l *LinkedList) List() {
	current := l.Head

	fmt.Println()
	for current != nil {
		fmt.Printf("Value : %d \t", current.Value)
		current = current.Next
	}
}

// Inplace edit of the linkedList
// With reversal of Nodes
func (l *LinkedList) Reverse() {
	// One approach is to create another linnked List
	// Keep Preprending the values in the order of
	// Current List will give us required result

	// Another approach is acting on the links
	// which ensures that updates happen in-place
	current := l.Head
	var previous *Node

	for current != nil {
		next := current.Next
		if next == nil {
			l.Head = current
		}
		current.Next = previous
		previous = current
		current = next
	}
}

func (l *LinkedList) RecursiveReverse(previous, current *Node) {
	if current == nil {
		return
	}

	next := current.Next
	if next == nil {
		l.Head = current
	}
	current.Next = previous
	previous = current

	l.RecursiveReverse(previous, next)
}

func main() {
	ll := NewLinkedList()

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)

	ll.List()
	ll.RecursiveReverse(nil, ll.Head)
	ll.List()
}
