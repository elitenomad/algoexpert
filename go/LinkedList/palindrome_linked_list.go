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

func Reverse(o LinkedList) LinkedList {
	var l LinkedList
	l = o

	current := l.Head
	var previous *Node

	for current != nil { // O(N)
		next := current.Next
		if next == nil {
			l.Head = current
		}
		current.Next = previous
		previous = current
		current = next
	}

	return l
}

func (l *LinkedList) Palindrome() bool {
	reverse := Reverse(*l)
	current := l.Head
	reverse_current := reverse.Head

	if reverse_current.Value != current.Value {
		return false
	}

	for current != nil {
		if current.Value != reverse_current.Value {
			return false
		}

		current = current.Next
		reverse_current = reverse_current.Next
	}

	return true
}

func (l *LinkedList) PalindromeII() bool {
	// Find middle element
	// For odd nodes its the middle one
	// For even nodes its the lowest of two
	slow, fast := l.Head, l.Head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the elements after middle element
	middle := slow
	current := middle.Next
	var previous *Node
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	// Compare elements of first half of the linked list
	// with later half.
	current = l.Head
	for current.Next != slow {
		if current.Value != middle.Next.Value {
			return false
		}

		current = current.Next
		middle = middle.Next
	}

	return true
}

func main() {
	ll := NewLinkedList()

	ll.Append(1)
	ll.Append(2)
	ll.Append(2)
	ll.Append(3)

	fmt.Println(ll.PalindromeII())
}
