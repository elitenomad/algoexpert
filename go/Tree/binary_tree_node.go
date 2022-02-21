package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func New() *Node {
	return &Node{}
}

func Serialize(root *Node) []int {
	result := []int{}
	SerializeHelper(root, &result)

	return result
}

func SerializeHelper(root *Node, result *[]int) {
	if root == nil {
		*result = append(*result, -1)
		return
	}

	*result = append(*result, root.Key)
	SerializeHelper(root.Left, result)
	SerializeHelper(root.Right, result)
}

func Deserialize(input []int) *Node {
	index := 0
	root := &Node{}
	DeserializeHelper(input, &index, root)
	return root
}

func DeserializeHelper(input []int, index *int, root *Node) *Node {
	if *index == len(input) {
		return nil
	}

	val := input[*index]
	(*index) += 1

	if val == -1 {
		return nil
	}

	root.Key = val
	root.Left = DeserializeHelper(input, index, root)
	root.Right = DeserializeHelper(input, index, root)
	return root
}

func main() {
	node := New()
	node.Key = 10
	node.Left = New()
	node.Left.Key = 5
	node.Right = New()
	node.Right.Key = 20
	fmt.Println(node)
}
