package main

import (
	"fmt"
	"math"
)

func MaxAllocationOfPages(input []int, k int) int {
	// Find the sum and max elements
	high := 0
	low := math.MinInt32
	result := -1

	for i := 0; i < len(input); i++ {
		low = max(low, input[i])
		high += input[i]
	}
	fmt.Println(low, high)

	for low <= high {
		mid := (low + high) / 2

		if isValid(input, len(input), k, mid) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// Here we have sum
	return result
}

func isValid(arr []int, n int, k int, mid int) bool {
	tempK := 1
	sum := 0

	for i := 0; i < len(arr); i++ {
		if sum+arr[i] > mid {
			tempK += 1
			sum = arr[i]
		} else {
			sum += arr[i]
		}
	}
	return tempK <= k
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := []int{10, 20, 30, 40}
	k := 2

	fmt.Println(MaxAllocationOfPages(a, k))
}
