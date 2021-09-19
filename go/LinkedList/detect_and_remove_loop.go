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

func (l *LinkedList) CreateCycle(to, from int) {
	current := l.Head
	var cycleTo *Node
	var cycleFrom *Node

	for current != nil {
		if current.Value == to {
			cycleTo = current
		}

		if current.Value == from {
			cycleFrom = current
		}
		current = current.Next
	}

	cycleFrom.Next = cycleTo
}

func (l *LinkedList) List() {
	current := l.Head

	fmt.Println()
	for current != nil {
		fmt.Printf("Value : %d \t", current.Value)
		current = current.Next
	}
}

//Time Complexity O(N)
// Space Complexity O(N)
func (l *LinkedList) DetectAndRemoveLoop() {
	h := map[int]bool{}
	current := l.Head

	// Use case 1 - If Node is Nil there no cycle
	if current == nil {
		return
	}

	// Use case 2 - If there is only one Node in the List
	if current.Next == nil {
		return
	}

	// Use case 3 - If Linked List ends with last Node
	var previous *Node
	for current != nil {
		h[current.Value] = true // Store the value

		previous = current
		current = current.Next
		if current != nil && h[current.Value] {
			previous.Next = nil
			current = previous.Next
		}
	}
}

func (l *LinkedList) DetectAndRemoveLoopUsingFYOD() {
	slow, fast := l.Head, l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if slow != fast {
		return
	}

	slow = l.Head
	for slow.Next != fast.Next {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = nil
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

	ll.CreateCycle(15, 20)
	ll.DetectAndRemoveLoopUsingFYOD()
	ll.List()
}
