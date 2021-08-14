package main

import "fmt"

// Goal is we try to Insert the elements by row and
// Increment the count when insert returns true
const (
	BINARY_SIZE = 2
)

type TrieNode struct {
	children  [BINARY_SIZE]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func InitializeTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(in []int) bool {
	flag := false
	currentNode := t.root

	for i := 0; i < len(in); i++ {
		index := 0
		if in[i] == 1 {
			index = 1
		}

		if currentNode.children[index] == nil {
			currentNode.children[index] = &TrieNode{}
			flag = true
		}

		currentNode = currentNode.children[index]
	}

	return flag
}

func CountDistinctRowsInBM(input [][]int) int {
	count := 0
	t := InitializeTrie()

	for i := 0; i < len(input); i++ {
		if t.Insert(input[i]) {
			count++
		}
	}
	return count
}

func main() {
	a := [][]int{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 1},
	}
	fmt.Println(CountDistinctRowsInBM(a))
}
