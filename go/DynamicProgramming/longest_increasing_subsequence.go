package main

import "fmt"

// We are planning to find the longest increasing
// Sequence until that index
func LongestIncreasingSubsequence(input []int) int {
	dp := make([]int, len(input))
	dp[0] = 1

	for i := 1; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if input[i] > input[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	result := 0

	for i := 0; i < len(dp); i++ {
		result = max(result, dp[i])
	}

	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := []int{3, 4, 2, 8, 10, 5, 1}
	fmt.Println(LongestIncreasingSubsequence(a))
}
