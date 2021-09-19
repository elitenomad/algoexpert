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
func (l *LinkedList) DetectLoop() bool {
	h := map[int]bool{}
	current := l.Head

	// Use case 1 - If Node is Nil there no cycle
	if current == nil {
		return false
	}

	// Use case 2 - If there is only one Node in the List
	if current.Next == nil {
		return false
	}

	// Use case 3 - If Linked List ends with last Node
	for current != nil {
		h[current.Value] = true // Store the value

		current = current.Next
		if current != nil && h[current.Value] {
			return true
		}
	}

	return false
}

func (l *LinkedList) DetectLoopUsingFyodAlgorithm() bool {
	// Maintain two pointers
	// One Pointer is slow pointer
	// Another Pointer is a fast Pointer
	current, first, second := l.Head, l.Head, l.Head

	// for second != nil and second.next != nil
	for current != nil {
		first = first.Next
		if second == nil {
			return false
		}
		second = second.Next.Next

		current = current.Next
		if first == second {
			return true
		}
	}

	return false
}

//Time Complexity O(N)
// Space Complexity O(1)
func (l *LinkedList) DetectLoopSpaceEfficient() bool {
	current := l.Head

	// Use case 1 - If Node is Nil there no cycle
	if current == nil {
		return false
	}

	// Use case 2 - If there is only one Node in the List
	if current.Next == nil {
		return false
	}

	// Use case 3 - If Linked List ends with last Node
	for current != nil {
		current.Visited = true // Store the value

		current = current.Next
		if current != nil && current.Visited {
			return true
		}
	}

	return false
}

// Another Navie approach could be
// For each iteration go through the entire LL until that point to verify if its Next exists
// Time complexity would be O(N^2)

func main() {
	ll := NewLinkedList()

	ll.Append(10)
	ll.Append(15)
	ll.Append(12)
	ll.Append(20)

	// ll.CreateCycle(15, 20)
	fmt.Println(ll.DetectLoop())
	fmt.Println(ll.DetectLoopSpaceEfficient())
	fmt.Println(ll.DetectLoopUsingFyodAlgorithm())
}
