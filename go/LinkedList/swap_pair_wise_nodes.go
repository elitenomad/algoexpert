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

func (l *LinkedList) PairwiseSwapNodes() { // Swapping the values
	current := l.Head

	for current != nil && current.Next != nil {
		temp := current.Value
		current.Value = current.Next.Value
		current.Next.Value = temp

		current = current.Next.Next
	}
}

func (l *LinkedList) PairwiseSwapNodesUsingRefs() { // Swapping the references
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	current := l.Head.Next.Next // Jump two places

	// Swap References of first two nodes
	previous := l.Head
	l.Head = l.Head.Next
	l.Head.Next = previous

	for current != nil && current.Next != nil {
		previous.Next = current.Next
		previous = current
		next := current.Next.Next
		current.Next.Next = current
		current = next
	}
	previous.Next = current
}

func main() {
	ll := NewLinkedList()

	ll.Append(5)
	ll.Append(8)
	ll.Append(7)
	ll.Append(10)
	ll.Append(12)
	// ll.Append(15)

	ll.List()
	ll.PairwiseSwapNodesUsingRefs()
	ll.List()
}
