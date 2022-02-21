package main

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func FindPath(root *Node, path []int, n int) bool {
	if root == nil {
		return false
	}

	path = append(path, root.Key)

	if root.Key == n {
		return true
	}

	if FindPath(root.Left, path, n) || FindPath(root.Right, path, n) {
		return true
	}

	return false
}

func LowestCommonAncestor(root *Node, n1 int, n2 int) *Node {
	pathOne := []int{}
	pathTwo := []int{}

	if FindPath(root, pathOne, n1) == false || FindPath(root, pathTwo, n2) == false {
		return nil
	}

	for i := 0; i < len(pathOne) && i < len(pathTwo); i++ {
		if pathOne[i] != pathTwo[i] {
			return &Node{Key: pathOne[i]}
		}
	}

	return nil
}

func main() {

}
