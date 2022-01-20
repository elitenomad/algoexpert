package main

import (
	"container/list"
	"fmt"
)

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

func (l *LinkedList) Reverse() *LinkedListNode {
	current := l.Head
	var previous *LinkedListNode

	for current != nil {
		next := current.Next
		if current.Next == nil { // Make sure Head is present in the LinkedList
			l.Head = current
		}
		current.Next = previous
		previous = current
		current = next
	}

	return previous
}

type Position struct {
	x int
	y int
}

func main() {
	// linkedList := New()
	// linkedList.append(1)
	// linkedList.append(2)
	// linkedList.append(3)
	// linkedList.append(4)
	// linkedList.Length()
	// linkedList.List()
	// linkedList.prepend(0)
	// linkedList.List()
	// linkedList.Length()
	// fmt.Println(linkedList.find(5))
	// linkedList.deleteHead()
	// linkedList.List()
	// linkedList.Length()
	// linkedList.deleteTail()
	// linkedList.List()
	// linkedList.Length()
	// linkedList.delete(3)
	// fmt.Println(linkedList.Reverse().Next.Next.Next)
	// linkedList.List()
	// linkedList.Length()

	fmt.Println("Go Linked Lists Tutorial")

	mylist := list.New()

	pos_1 := Position{x: 1, y: 1}
	pos_2 := Position{x: 10, y: 10}
	mylist.PushBack(pos_1)
	mylist.PushFront(pos_2)

	for element := mylist.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		fmt.Println(element.Value.(Position).x)
	}
}
