package main

import (
	"fmt"
	"math"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func CeilOnLeftOfArray(input []int) []int {
	result := make([]int, len(input))
	result[0] = -1

	for i := 1; i < len(input); i++ {
		diff := math.MaxInt32
		for j := 0; j < i; j++ {
			if input[j] >= input[i] {
				diff = min(diff, input[j])
			}
		}

		if diff == math.MaxInt32 {
			result[i] = -1
		} else {
			result[i] = diff
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func CeilOnLeftOfArrayII(input []int) {
	tree := rbt.NewWithIntComparator()
	tree.Put(input[0], input[0])
	fmt.Println(-1)
	for i := 1; i < len(input); i++ {
		n, found := tree.Ceiling(input[i])
		if !found {
			fmt.Println(-1)
		} else {
			fmt.Println(n)
		}

		tree.Put(input[i], input[i])
	}
}

func main() {
	a := []int{2, 8, 30, 15, 25, 12}
	fmt.Println(CeilOnLeftOfArray(a))
	CeilOnLeftOfArrayII(a)
}
