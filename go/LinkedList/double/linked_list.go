package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node // Do we really need Tail ?
	Size int
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Size: 0,
	}
}

func (l *DoublyLinkedList) Append(value int) {
	node := NewNode(value)

	if l.Head == nil {
		// Set Prev and Next as nil
		node.Prev = nil
		node.Next = nil

		// Set list head to node created
		l.Head = node
		return
	}

	// If you maintain Tail as well
	// Only thin you we have to change is the tail
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
	current.Next.Prev = current
}

func (l *DoublyLinkedList) Prepend(value int) {
	node := NewNode(value)

	if l.Head == nil {
		// Set Prev and Next as nil
		node.Prev = nil
		node.Next = nil

		// Set list head to node created
		l.Head = node
		return
	}

	// collect current Head
	previousHead := l.Head
	previousHead.Prev = node
	l.Head = node
	l.Head.Next = previousHead
	// l.Head.Prev will be nil
}

// This reverse requires O(N) of auxillary space
func (l *DoublyLinkedList) Reverse() *DoublyLinkedList {
	ddl := &DoublyLinkedList{}
	current := l.Head

	// If linked list is nil or just have only one element
	// Do nothing and just return
	if current == nil || current.Next == nil {
		return ddl
	}

	// var prev *Node
	for current != nil {
		ddl.Prepend(current.Value) // O(1)
		// Move to the next element
		current = current.Next
	}

	return ddl
}

func (l *DoublyLinkedList) DeleteHead() {
	current := l.Head

	// When linked list is empty
	// Or have one node
	// Empty the head and return
	if current == nil || current.Next == nil {
		l.Head = nil

		return
	}

	// Move current head to current Head's Next
	l.Head = current.Next

	// Make the current Head prev to nil
	l.Head.Prev = nil
}

func (l *DoublyLinkedList) DeleteTail() {
	current := l.Head

	// When linked list is empty
	// Or have one node
	// Empty the head and return
	if current == nil || current.Next == nil {
		l.Head = nil

		return
	}

	for current.Next != nil {
		current = current.Next
	}

	current.Prev.Next = nil
	current.Prev = nil
}

func (l *DoublyLinkedList) Print() {
	current := l.Head

	for current != nil {
		fmt.Println("Prev: ", current.Prev, " Next:", current.Next, " Value: ", current.Value)
		current = current.Next
	}
}

func main() {
	/*
		Goal is to manually create a Doubly Linked list with values
		10, 20 and 30
	*/
	ll := NewDoublyLinkedList()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Prepend(5)
	ll.Print()
	fmt.Println("")
	ll.Reverse().Print()

	// ll.DeleteHead()
	// ll.Print()
	// fmt.Println()

	// ll.DeleteTail()
	// ll.Print()
}
