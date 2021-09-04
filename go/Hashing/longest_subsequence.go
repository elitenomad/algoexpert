package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Naive(input []int) int { //O(NLOGN)
	sort.Ints(input)
	// fmt.Println(input)
	result := 1
	current := 1
	for i := 0; i < len(input)-1; i++ {
		if input[i+1]-input[i] == 1 {
			current += 1
		} else {
			result = max(result, current)
			current = 1
		}
	}

	return max(result, current)
}

func Efficient(input []int) int {
	result := 1
	h := map[int]bool{}

	for i := 0; i < len(input); i++ { //O(N)
		h[input[i]] = true
	}
	// We will remove all the duplicate elements from array

	for k, _ := range h {
		current := 1
		if _, found := h[k-1]; !found {
			for h[k+current] {
				current += 1
			}

			result = max(result, current)
		}
	}

	return result
}

func main() {
	a := []int{1, 9, 3, 4, 2, 20}
	fmt.Println(Naive(a))
	fmt.Println(Efficient(a))

	a = []int{8, 20, 7, 30}
	fmt.Println(Naive(a))
	fmt.Println(Efficient(a))

	a = []int{20, 30, 40}
	fmt.Println(Naive(a))
	fmt.Println(Efficient(a))
}
