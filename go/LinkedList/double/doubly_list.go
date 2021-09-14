/*
	Doubly Linked List maintaining Head and Tail
*/
package main

type Node struct {
	Value      int
	Prev, Next *Node
}

type DoublyLinkedList struct {
	Head, Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (ll *DoublyLinkedList) SetHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node

		return
	}

	ll.InsertBefore(ll.Head, node)
}

func (ll *DoublyLinkedList) SetTail(node *Node) {
	if ll.Tail == nil {
		ll.Head = node
		ll.Tail = node

		return
	}

	ll.InsertAfter(ll.Tail, node)
}

func (ll *DoublyLinkedList) InsertBefore(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	// Remove Node from List
	ll.Remove(nodeToInsert)

	nodeToInsert.Prev = node.Prev
	nodeToInsert.Next = node

	if node.Prev == nil {
		ll.Head = nodeToInsert
	} else {
		node.Prev.Next = nodeToInsert
	}

	node.Prev = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAfter(node, nodeToInsert *Node) {
	if nodeToInsert == ll.Head && nodeToInsert == ll.Tail {
		return
	}

	ll.Remove(nodeToInsert)

	// Set Next prev for nodeToInsert
	nodeToInsert.Prev = node
	nodeToInsert.Next = node.Next

	// Set Next and Prev for node
	if node.Next == nil {
		ll.Tail = nodeToInsert
	} else {
		node.Next.Prev = nodeToInsert
	}

	node.Next = nodeToInsert
}

func (ll *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *Node) {
	if position == 1 {
		ll.SetHead(nodeToInsert)
		return
	}

	node := ll.Head
	currentPosition := 1
	for node != nil && currentPosition != position {
		node = node.Next
		currentPosition += 1
	}

	if node != nil {
		ll.InsertBefore(node, nodeToInsert)
	} else {
		ll.SetTail(nodeToInsert)
	}
}

func (ll *DoublyLinkedList) RemoveNodesWithValue(value int) {
	node := ll.Head

	for node != nil {
		nodeToRemove := node // Maintain another variable for deletion
		node = node.Next
		if nodeToRemove.Value == value {
			ll.Remove(nodeToRemove)
		}
	}
}

func (ll *DoublyLinkedList) Remove(node *Node) {
	// What if node is Head
	if ll.Head == node {
		ll.Head = ll.Head.Next
	}

	// If node is Tail
	if node == ll.Tail {
		ll.Tail = ll.Tail.Prev
	}

	ll.removeNodeBindings(node)
}

func (ll *DoublyLinkedList) removeNodeBindings(node *Node) {
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	node.Prev = nil
	node.Next = nil
}

func (ll *DoublyLinkedList) ContainsNodeWithValue(value int) bool {
	current := ll.Head

	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}

	return false
}
