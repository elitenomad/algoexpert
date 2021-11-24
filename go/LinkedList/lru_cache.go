package main

import "fmt"

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct { // DLL
	Head *Node
	Tail *Node
	Size int
}

type LRUCache struct {
	Capacity int
	Size     int
	Data     map[int]*Node
	List     *LinkedList
}

func New(cap int) *LRUCache {
	return &LRUCache{
		Capacity: cap,
		List:     &LinkedList{},
	}
}

func (l *LinkedList) List() {
	current := l.Head

	fmt.Println()
	for current != nil {
		fmt.Printf("Value : %d \t", current.Value)
		current = current.Next
	}
}

func (l *LinkedList) DeleteNode(node *Node) {
	node.Prev.Next = node.Next.Next
	node.Next.Prev = node.Prev
}

func (l *LinkedList) AddToHead(node *Node) {
	l.Head.Prev = node
	node.Next = l.Head
	l.Head = node // decide on head
}

func (c *LRUCache) Get(key, val int) int {
	// Get function should not only get the value
	// But should move the node found to the MRU
	if node, found := c.Data[key]; found {
		result := node.Value
		c.List.DeleteNode(node)
		c.List.AddToHead(node)
		c.Data[key] = node // Not required but ensuring
		return result
	}

	return -1
}

func (c *LRUCache) Set(key, val int) {
	// First find the value by the Key
	// 	If it exists, replace the value
	//	Delete Node
	// 	Add it to Head

	// If the Key do not exist
	// Verify if we reached the capacity
	//	If we reached Capacity delete tail from list and map
	// 	Add the node to the head
	// 	If we didnt reach the capacity
	//	Increase the size and add the node to head.
	if node, found := c.Data[key]; found {
		node.Value = val
		c.List.DeleteNode(node)
		c.List.AddToHead(node)
	} else {
		node := &Node{Key: key, Value: val}
		c.Data[key] = node
		if c.Size < c.Capacity {
			c.Size += 1
			c.List.AddToHead(node)
		} else {
			delete(c.Data, c.List.Tail.Key)
			c.List.DeleteNode(c.List.Tail)
			c.List.AddToHead(node)
		}
	}
}

func main() {
	// ll := New(5)

}
