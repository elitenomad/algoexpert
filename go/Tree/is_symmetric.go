package main

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func New() *Node {
	return &Node{}
}

func isSymmetric(root *Node) bool {
	return isMirror(root, root)
}

func isMirror(n1 *Node, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	return (n1.Key == n2.Key) && isMirror(n1.Right, n2.Left) && isMirror(n1.Left, n2.Right)

}

func main() {

}
