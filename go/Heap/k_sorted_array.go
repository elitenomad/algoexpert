package main

import (
	"fmt"
)

// Sort the array
// An element at index i will be present b/w i-k and i+k
func KSortedArray(input []int, k int) []int {
	// Implementation of Priority Queue is something to look for
	// Lets assume we have PQ implementation set
	pq := NewPriorityQueue() // Integers
	for i := 0; i <= k; i++ {
		pq.Add(input[i])
	}

	index := 0
	for i := k + 1; i < len(input); i++ {
		input[index] = pq.Poll()
		index += 1
		pq.Add(input[i])
	}

	for !pq.IsEmpty() {
		input[index] = pq.Poll()
		index += 1
	}

	return input
}

func main() {
	input := []int{15, 20, 4, 15, 50}
	fmt.Println(KSortedArray(input, 2))
}
