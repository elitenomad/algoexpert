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

// Time complexity O(N)
// Space Complexity O(N/K)
func (l *LinkedList) RecursiveReverseInGroups(head *Node, k int) *Node {
	current := head
	var next *Node
	var previous *Node
	count := 0

	for current != nil && count < k {
		next = current.Next
		current.Next = previous
		previous = current
		current = next

		count += 1
	}

	if next != nil {
		head.Next = l.RecursiveReverseInGroups(next, k)
	}

	l.Head = previous // Sets the head
	return l.Head
}

func (l *LinkedList) ReverseInGroups(head *Node, k int) *Node {
	current := head
	var next *Node
	var previous *Node
	var previousFirst *Node
	isFirstPass := true

	for current != nil { // Loop Until the end of LinkedList
		count := 0
		first := current
		// Loop multiple times < K until Linked list is complete
		// Below loop is responsible for counter increment
		for current != nil && count < k {
			next = current.Next
			current.Next = previous
			previous = current
			current = next

			count += 1
		}
		if isFirstPass {
			l.Head = previous
			isFirstPass = false
		} else {
			previousFirst.Next = previous
		}

		previousFirst = first
	}

	return l.Head
}

func main() {
	ll := NewLinkedList()

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)
	ll.Append(50)
	ll.Append(60)

	ll.List()
	ll.RecursiveReverseInGroups(ll.Head, 2)
	ll.List()

	ll.RecursiveReverseInGroups(ll.Head, 2)
	ll.List()
}
