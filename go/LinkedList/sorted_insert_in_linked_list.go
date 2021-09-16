package main

import "fmt"

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
	Size int
}

func New() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *LinkedList) prepend(value int) {
	newNode := LinkedListNode{
		Value: value,
		Next:  l.Head,
	}

	l.Head = &newNode

	if l.Tail == nil {
		l.Tail = &newNode
	}

	l.Size += 1
}

func (l *LinkedList) append(value int) {
	newNode := LinkedListNode{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = &newNode
		l.Tail = &newNode
	}

	l.Tail.Next = &newNode
	l.Tail = &newNode
	l.Size += 1
}

func (l *LinkedList) delete(value int) *LinkedListNode {
	if l.Head == nil { // A case where linkedList itself is nil
		return nil
	}

	var deletedNode *LinkedListNode
	// If value is equal to head value
	for l.Head != nil && l.Head.Value == value {
		deletedNode = l.Head
		l.Head = l.Head.Next
		l.Size -= 1
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			deletedNode = currentNode.Next
			currentNode.Next = deletedNode.Next
			l.Size -= 1
		} else {
			currentNode = currentNode.Next
		}
	}

	if l.Tail.Value == value {
		l.Tail = currentNode
		l.Size -= 1
	}

	return deletedNode
}

func (l *LinkedList) deleteHead() {
	l.Head = l.Head.Next
	l.Size -= 1
}

func (l *LinkedList) deleteTail() *LinkedListNode {
	// Use case where linkedlist is nil
	if l.Head == nil {
		return nil
	}

	// Tail to be deleted
	currentTail := l.Tail

	// Start from the head
	currentNode := l.Head

	for currentNode.Next != nil {
		// Two things we need to do
		// when currentNode.Next is equal to currentTail, set that to nil
		// currentNode to Tail
		if currentNode.Next == currentTail {
			currentNode.Next = nil
			l.Tail = currentNode
			l.Size -= 1
		} else {
			currentNode = currentNode.Next
		}
	}

	return currentTail
}

func (l *LinkedList) find(value int) *LinkedListNode {
	currentNode := l.Head

	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}

	return nil
}

func (l *LinkedList) List() {
	currentNode := l.Head

	for currentNode != nil {
		fmt.Printf("Value: %d \t", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func (l *LinkedList) Length() {
	fmt.Println("Size ::", l.Size)
}

func (l *LinkedList) SortedInsertOfItem(value int) {
	head := l.Head
	node := &LinkedListNode{Value: value}

	if head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	if head.Value >= value {
		node.Next = head
		l.Head = node

		return
	}

	if value >= l.Tail.Value {
		l.Tail.Next = node
		l.Tail = node

		return
	}

	current := l.Head
	for current.Next.Value <= value {
		current = current.Next
	}

	node.Next = current.Next
	current.Next = node
}

func main() {
	linkedList := New()
	linkedList.append(10)
	linkedList.append(20)
	linkedList.append(30)

	linkedList.SortedInsertOfItem(5)
	linkedList.SortedInsertOfItem(15)
	linkedList.SortedInsertOfItem(35)

	linkedList.List()
}
