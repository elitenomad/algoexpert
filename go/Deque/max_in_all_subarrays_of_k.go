package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

// Brute Force
func MaxInAllSubArraysOfK(input []int, k int) []int {
	result := make([]int, len(input)-k+1)

	for i := 0; i < len(input)-k+1; i++ {
		for j := i + 1; j < i+3; j++ {
			if input[i] >= input[j] {
				result[i] = max(result[i], input[i])
			} else {
				result[i] = max(result[i], input[j])
			}
		}
	}

	return result
}

// Efficient approach implementation could be implemente4d
// Using Deque Data structure

func MaxOfEachSubArrayofSizeK(input []int, k int) {
	var deque deque.Deque

	for i := 0; i < k; i++ {
		for (deque.Len() > 0) && input[i] >= input[deque.Back().(int)] {
			deque.PopBack()
		}

		deque.PushBack(i)
	}

	for i := k; i < len(input); i++ {

		fmt.Println(input[deque.Front().(int)])

		for deque.Len() > 0 && deque.Front().(int) <= i-k {
			deque.PopFront()
			fmt.Println(deque)
		}

		for deque.Len() > 0 && input[i] >= input[deque.Back().(int)] {
			deque.PopBack()
		}

		deque.PushBack(i)
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := []int{10, 8, 5, 12, 15, 7, 6}
	k := 3
	MaxOfEachSubArrayofSizeK(a, k)

	// var q deque.Deque
	// q.PushBack("foo")
	// q.PushBack("bar")
	// q.PushBack("baz")

	// fmt.Println(q.Len())   // Prints: 3
	// fmt.Println(q.Front()) // Prints: foo
	// fmt.Println(q.Back())  // Prints: baz

	// q.PopFront() // remove "foo"
	// q.PopBack()  // remove "baz"

	// q.PushFront("hello")
	// q.PushBack("world")

	// // Consume deque and print elements.
	// for q.Len() != 0 {
	// 	fmt.Println(q.PopFront())
	// }
}
