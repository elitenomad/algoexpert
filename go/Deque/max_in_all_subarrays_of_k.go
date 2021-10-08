package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
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
	d := deque.New()

	for i := 0; i < k; i++ {
		for !deque.IsEmpty() && deque.GetRear() >= 0 && input[i] >= input[deque.GetRear()] {
			deque.DeleteRear()
		}
		deque.InsertRear(i)
	}

	for i := k; i < len(input); i++ {
		if deque.GetRear() >= 0 {
			fmt.Println(input[deque.GetRear()])
		}

		for !deque.IsEmpty() && deque.GetRear() <= i-k {
			deque.DeleteFront()
		}

		for !deque.IsEmpty() && deque.GetRear() >= 0 && input[i] >= input[deque.GetRear()] {
			deque.DeleteRear()
		}

		deque.InsertRear(i)
	}

	if deque.GetRear() >= 0 {
		fmt.Println(input[deque.GetRear()])
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
}
