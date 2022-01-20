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

func lengthOfLIS(nums []int) int {
	tail := []int{}

	tail = append(tail, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > tail[len(tail)-1] {
			tail = append(tail, nums[i])
		} else {
			c := ceilIdx(tail, 0, len(tail)-1, nums[i])
			tail[c] = nums[i]
		}
	}

	return len(tail)
}

func ceilIdx(tail []int, start int, end int, x int) int {
	for start < end {
		m := start + (end-start)/2

		if tail[m] >= x {
			end = m
		} else {
			start = m + 1
		}
	}

	return end
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
