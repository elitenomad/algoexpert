package main

type Node struct {
	Value int
	Next  *Node
}

// First approach i thought of is
// Create an Empty Linked List
// Create an Append method which maintains Head and Tail of LinkedList created
// Start the loop with I and J as two LinkList Heads
// Loop Until one of them is nil (If they are of same size both will be nil at the same time)
// Compare values at each iteration and append the small value to the newLy created Linked List
// Once the Loop is exited
// Ensure any left over elements of either of the Given input linked lists are looped through and
// Appended to the Newly created linked List
// Time Complexity is O(Max(M,N)) Space Complexity O(M+N)
func MergeTwoSortedLinkedLists(a, b *Node) *Node {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	// Initialise two nodes
	var head *Node
	var tail *Node

	// Compare first nodes and initialise head and tail
	if a.Value <= b.Value {
		head = a
		tail = a
		a = a.Next
	} else {
		head = b
		tail = b
		b = b.Next
	}

	for a != nil && b != nil {
		if a.Value <= b.Value {
			tail.Next = a
			tail = a
			a = a.Next
		} else {
			tail.Next = b
			tail = b
			b = b.Next
		}
	}
	if a == nil {
		tail.Next = b
	} else {
		tail.Next = a
	}

	return head
}

func main() {

}
