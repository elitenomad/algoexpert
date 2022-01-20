type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) DeleteNode(node *Node) {
	if node == nil || (node.Next == nil && node.Prev == nil) {
		l.Head = nil
		l.Tail = nil

		return
	}

	// Tail
	if node.Next == nil {
		node.Prev.Next = nil
		l.Tail = node.Prev
		return
	}

	//Head
	if node.Prev == nil {
		node.Next.Prev = nil
		l.Head = node.Next
		return
	}

	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next.Next
}

func (l *LinkedList) AddToHead(node *Node) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node

		return
	}

	l.Head.Prev = node
	node.Next = l.Head
	l.Head = node

}

type LRUCache struct {
	Capacity int
	Size     int
	Data     map[int]*Node
	List     *LinkedList
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		List:     &LinkedList{},
		Data:     make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, found := this.Data[key]; found {
		// Delete the node
		// Add it to the head
		// Add it to hash again (not required but to ensure the value is not lost)
		result := node.Value
		this.List.DeleteNode(node)
		this.List.AddToHead(node)
		this.Data[key] = node
		return result
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, found := this.Data[key]; found {
		node.Value = value // Just in case the value is changed in the next RUN
		this.List.DeleteNode(node)
		this.List.AddToHead(node)
	} else {
		node := &Node{Key: key, Value: value}
		this.Data[key] = node

		if this.Size < this.Capacity {
			this.Size += 1
			this.List.AddToHead(node)
		} else {
			delete(this.Data, this.List.Tail.Key)
			this.List.DeleteNode(this.List.Tail)
			this.List.AddToHead(node)
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */