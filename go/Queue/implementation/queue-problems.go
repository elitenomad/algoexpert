package main

/*
	Implement stack using Queue by making pop operation costly
	Implement stack using only one queue (we use recursion call stack here)
	Implement Queue using stack
		- By making enqueu operation costly
		- By making dequeu operation costly
		- By using one stack (And one recursion call stack)
*/

// In order to implement stack using Queue we need Queue implementation
type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
}
type Queue struct {
	Element *LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		Element: &LinkedList{},
	}
}

func (q *Queue) Enqueue(element int) {
	node := &Node{Value: element}

	if q.Element.Head == nil {
		q.Element.Head = node
		q.Element.Tail = node
		return
	}

	q.Element.Tail.Next = node
	q.Element.Tail = node
}

func (q *Queue) Dequeue() *Node {
	nodeToBeRemoved := q.Element.Head
	q.Element.Head = q.Element.Head.Next

	nodeToBeRemoved.Next = nil
	return nodeToBeRemoved
}

func (q *Queue) IsEmpty() bool {
	return q.Element.Head == nil
}

// Implement stack now given we have to use just queue
type Stack struct {
	Q1   *Queue
	Q2   *Queue // Below implementation using only One Queue
	Size int
}

func NewStack() *Stack {
	return &Stack{
		Q1: NewQueue(),
		Q2: NewQueue(),
	}
}

func (s *Stack) Push(element int) {
	s.Q1.Enqueue(element)

}

// Another version of implementing stack is
// use two queues
// In that case we will have constly Push operation
// and O(1) pop operation because in that inscenario
// We have to just replace the head
func (s *Stack) Pop() int { //O(N) Implementation with Pop operation costly
	elementToBePopped := s.Q1.Element.Tail
	current := s.Q1.Element.Head

	for current.Next.Next != nil {
		current = current.Next
	}

	s.Q1.Element.Tail = current
	current.Next = nil

	return elementToBePopped.Value
}

func (s *Stack) IsEmpty() bool {
	return s.Size <= 0
}

func main() {
	// queue := NewQueue()
	// queue.Enqueue(1)
	// queue.Enqueue(2)
	// queue.Enqueue(3)
	// queue.Enqueue(4)
	// fmt.Println()
	// fmt.Println(queue.IsEmpty())
	// fmt.Println(queue.Element.Head, queue.Element.Tail)
	// queue.Dequeue()
	// fmt.Println(queue.Element.Head, queue.Element.Tail)
	// queue.Dequeue()
	// queue.Dequeue()
	// fmt.Println(queue.Element.Head, queue.Element.Tail)

	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
}
