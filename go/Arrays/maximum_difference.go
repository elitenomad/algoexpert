package main

import (
	"fmt"
	"math"
)

/*
	Brute force approach
*/
func MaxDifference(input []int) float64 {
	// Large -ve value
	max := math.Inf(-10)

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input)-1; j++ {
			diff := (input[j] - input[i])
			max = math.Max(float64(max), float64(diff))
		}
	}

	return max
}

// Efficient Approach O(n)
func MaxDifferenceEff(input []int) int {
	maximum := input[1] - input[0]
	minVal := input[0]

	for i := 1; i < len(input); i++ {
		maximum = max(maximum, input[i]-minVal)
		minVal = min(minVal, input[i])
		// fmt.Println(maximum, minVal)
	}

	return maximum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{2, 3, 10, 6, 4, 8, 1}
	fmt.Println(MaxDifference(arr))
	fmt.Println(MaxDifferenceEff(arr))
}
