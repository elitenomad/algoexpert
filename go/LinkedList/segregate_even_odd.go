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

func (l *LinkedList) SegregateEvenAndOdd() *LinkedList {
	ll := NewLinkedList()
	current := l.Head

	for current != nil { // O(N)
		if current.Value%2 == 0 {
			ll.Append(current.Value) // Optimise by maintaining a Tail in the LinkedList
		}

		current = current.Next
	}

	current = l.Head
	for current != nil { // O(N)
		if current.Value%2 != 0 {
			ll.Append(current.Value)
		}

		current = current.Next
	}

	return ll
}

func (l *LinkedList) SegregateEvenAndOddImprovise() *Node {
	current := l.Head

	// Instead of Maintaining 4 nodes, we can try this with 2 Linked List
	// One for Even and another for Odd
	var es *Node
	var ee *Node
	var os *Node
	var oe *Node

	for current != nil { // O(N)
		if current.Value%2 == 0 {
			if es == nil {
				es = current
				ee = es
			} else {
				ee.Next = current
				ee = ee.Next
			}
		} else {
			if os == nil {
				os = current
				oe = os
			} else {
				oe.Next = current
				oe = oe.Next
			}
		}

		current = current.Next
	}

	ee.Next = os
	oe.Next = nil

	return es
}

func main() {
	ll := NewLinkedList()

	ll.Append(17)
	ll.Append(15)
	ll.Append(8)
	ll.Append(12)
	ll.Append(10)
	ll.Append(5)
	ll.Append(4)

	ll.List()
	// tl := ll.SegregateEvenAndOdd()
	// tl.List()

	node := ll.SegregateEvenAndOddImprovise()
	c := node

	fmt.Println()
	for c != nil {
		fmt.Printf("Value: %d\t", c.Value)
		c = c.Next
	}
}
