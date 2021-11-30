package main

import (
	"fmt"
	"math"
)

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

func CountDistinctRowsInBMII(input [][]int) int {
	if len(input) <= 0 {
		return 0
	}

	rows := len(input)
	countsOf10PerRow := map[int]bool{}

	for i := 0; i < rows; i++ {
		countsOf10PerRow[MultiplyInIncrementsOf10(input[i])] = true
	}

	count := 0
	for _, v := range countsOf10PerRow {
		_ = v
		count += 1
	}

	return count
}

func MultiplyInIncrementsOf10(row []int) int {
	j := len(row) - 1
	sum := 0

	for i := 0; i < len(row); i++ {
		sum += row[i] * PowersOf10(10, j)
		j--
	}

	return sum
}

func PowersOf10(a, b int) int {
	return int(math.Pow(float64(3), float64(5)))
}

func main() {
	a := [][]int{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 1},
	}
	fmt.Println(CountDistinctRowsInBM(a))
	fmt.Println(CountDistinctRowsInBMII(a))
}
