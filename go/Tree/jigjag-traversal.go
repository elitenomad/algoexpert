package main

import "fmt"

type Queue struct {
	Elements  []*TreeNode
	direction int
}

func NewQueue() *Queue {
	return &Queue{direction: 1}
}

func (q *Queue) Enqueue(val *TreeNode) {
	q.Elements = append(q.Elements, val)
}

func (q *Queue) Deque() *TreeNode {
	d := q.Elements[0]
	q.Elements = q.Elements[1:]
	return d
}

func (q *Queue) Size() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	// direction := 1 // + anti-clockwise, - clockwise
	queue := NewQueue()
	queue.Enqueue(root)
	result := [][]int{}

	for !queue.IsEmpty() {
		fmt.Println(queue.Elements)
		if queue.direction == 1 {
			temp := []int{}
			for i := 0; i < len(queue.Elements); i++ {
				temp = append(temp, queue.Elements[i].val)
			}

			result = append(result, temp)
		} else {
			temp := []int{}
			for i := len(queue.Elements) - 1; i >= 0; i-- {
				temp = append(temp, queue.Elements[i].val)
			}

			result = append(result, temp)
		}

		size := queue.Size()
		for i := 0; i < size; i++ {
			elem := queue.Deque()
			if elem.left != nil {
				queue.Enqueue(elem.left)
			}

			if elem.right != nil {
				queue.Enqueue(elem.right)
			}
		}

		if queue.direction == 1 {
			queue.direction = -1
		} else {
			queue.direction = 1
		}
	}

	return result
}

func main() {
	node := &TreeNode{}
	node.val = 1
	node.left = &TreeNode{}
	node.left.val = 2
	node.right = &TreeNode{}
	node.right.val = 3
	node.left.left = &TreeNode{}
	node.left.right = &TreeNode{}
	node.left.left.val = 4
	node.left.right.val = 24
	node.right.left = &TreeNode{}
	node.right.right = &TreeNode{}
	node.right.left.val = 6
	node.right.right.val = 7
	fmt.Println(zigzagLevelOrder(node))
}
