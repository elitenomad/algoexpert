package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func New() *Node {
	return &Node{}
}

//Implement a Queue
type Queue struct {
	Elements []*Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(node *Node) {
	q.Elements = append(q.Elements, node)
}

func (q *Queue) Dequeue() *Node {
	dequeuedNode := q.Elements[0]
	q.Elements = q.Elements[1:]
	return dequeuedNode
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) <= 0
}

// Complete

func GrandParentEvenValues(tree *Node) int {
	sum := 0
	GrandParentEvenValuesHelper(tree, &sum, nil, nil)
	return sum
}

func GrandParentEvenValuesHelper(tree *Node, sum *int, parent *Node, grand_parent *Node) {
	if tree == nil {
		return
	}

	if grand_parent != nil && grand_parent.Val%2 == 0 {
		*sum += tree.Val
	}

	GrandParentEvenValuesHelper(tree.Left, sum, tree, parent)
	GrandParentEvenValuesHelper(tree.Right, sum, tree, parent)
}

func main() {
	node := New()
	node.Val = 10
	node.Left = New()
	node.Left.Val = 5
	node.Right = New()
	node.Right.Val = 20
	node.Right.Left = New()
	node.Right.Right = New()
	node.Right.Left.Val = 15
	node.Right.Right.Val = 40
	fmt.Println(GrandParentEvenValues(node))
}
