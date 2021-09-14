package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type CircularLinkedList struct {
	Head *Node
	Size int
	// We can have a Tail here but
	// For simplicity sake ignoring it
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{
		Size: 0,
		Head: nil,
	}
}

func (l *CircularLinkedList) Append(value int) {
	newNode := &Node{
		Value: value,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Head.Next = newNode
		l.Size += 1

		return
	}

	current := l.Head
	for current.Next != l.Head {
		current = current.Next
	}

	current.Next = newNode
	newNode.Next = l.Head
	l.Size += 1
}

func (l *CircularLinkedList) Prepend(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Head.Next = newNode
		l.Size += 1

		return
	}

	current := l.Head
	for current.Next != l.Head {
		current = current.Next
	}
	current.Next = newNode
	newNode.Next = l.Head
	l.Head = newNode

	/*
		Maintaining Tail in your LinkedList Structure makes it O(1) operation.

		Another approach is to attach the node next to Current Head and then
		swap the values once the node insertion is complete.
	*/
}

func (l *CircularLinkedList) DeleteHead() *Node {
	head := l.Head

	if head == nil || head.Next == nil {
		l.Head = nil

		return l.Head
	}

	// Swap Head's value with Heads Next value
	// Swap Head's Next with Headx Next.Next
	l.Head.Value = l.Head.Next.Value
	l.Head.Next = l.Head.Next.Next

	return head
}

func (l *CircularLinkedList) DeleteKthNode(k int) {
	// 1st Node
	// Middle nodes
	// Last Node
}

func (l *CircularLinkedList) Print() {
	current := l.Head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	// Implementation 1 using the for loop
	// fmt.Println("Value: ", current.Value, " Next: ", current.Next)
	// for r := current.Next; r != l.Head; r = r.Next {
	// 	fmt.Println("Value: ", r.Value, " Next: ", r.Next)
	// }

	for current != nil {
		fmt.Println("Value: ", current.Value, " Next: ", current.Next)

		// Ensure once you reach the end. Stop it :)
		if current.Next == l.Head {
			break
		}
		current = current.Next
	}
}

func main() {
	/*
		Goal is to manually create a Doubly Linked list with values
		10, 20 and 30
	*/
	ll := NewCircularLinkedList()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Prepend(5)
	ll.DeleteHead()

	ll.Print()
	fmt.Println("")
}
