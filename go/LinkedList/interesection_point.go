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

func IntersectionPoint(n1, n2 *Node) int { // Time Complexity O(M+N) Space O(M)
	h := map[int]bool{}

	current := n1
	for current != nil {
		h[current.Value] = true
		current = current.Next
	}

	current = n2
	for current != nil {
		if h[current.Value] {
			return current.Value
		}
		current = current.Next
	}

	return -1
}

func IntersectionPointImprovise(n1, n2 *Node) {
	// Find the Length of n1
	// Find the Length of n2
	// Find the abs(len(n1) - len(n2))
	// Loop through n1 until the size of above absolute - Mark this as current
	// Start looping through both nodes simulataniously until they meet on the same node
	// Return the node value

	/*
		for curr1 != nil && curr2 != nil {
			if curr1 == curr2 {
				return curr1.Value
			}

			curr1 = curr1.Next
			curr2 = curr2.Next
		}
	*/
}

func main() {
	ll := NewLinkedList()

	ll.Append(5)
	ll.Append(8)
	ll.Append(7)
	ll.Append(10)
	ll.Append(12)
	ll.Append(15)

	ml := NewLinkedList()
	ml.Append(9)
	ml.Append(10)
	ml.Append(11)
	ml.Append(12)

	// ll.List()
	fmt.Println(IntersectionPoint(ll.Head, ml.Head))
}
