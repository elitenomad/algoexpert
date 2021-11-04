package main

import "fmt"

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func New() *Tree {
	return &Tree{}
}

// Pre Order Traversal
func Solution(T *Tree) int {
	if T == nil {
		return 0
	}

	cache := map[int]int{}
	return SolutionHelper(T, &cache)
}

func SolutionHelper(T *Tree, cache *map[int]int) int {

	if T == nil || (*cache)[T.X] > 0 {
		return len(*cache)
	}

	if _, found := (*cache)[T.X]; found {
		(*cache)[T.X] += 1
	} else {
		(*cache)[T.X] = 1
	}
	fmt.Println(*cache)
	maxOfBothSides := max(SolutionHelper(T.L, cache), SolutionHelper(T.R, cache))

	// Somehow we need to remove dupes from hash
	if _, found := (*cache)[T.X]; found {
		(*cache)[T.X] -= 1
	}

	// At one point the data is becoming 0,
	// so see if we can delete as o/p is dependent on length of hash
	if (*cache)[T.X] == 0 {
		delete(*cache, T.X)
	}

	return maxOfBothSides
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	node := New()
	node.X = 1
	// node.L = New()
	// node.L.X = 5
	node.R = New()
	node.R.X = 2
	// node.L.L = New()
	// node.L.R = New()
	// node.L.L.X = 1
	// node.L.R.X = 1
	node.R.L = New()
	node.R.R = New()
	node.R.L.X = 1
	node.R.R.X = 1
	node.R.R.L = New()
	node.R.R.L.X = 4

	fmt.Println(Solution(node))
}
