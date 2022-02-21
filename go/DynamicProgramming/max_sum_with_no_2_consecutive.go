package main

import "fmt"

func MaxSumWithNoTwoConsecutivesRec(input []int) int {
	return MaxSumWithNoTwoConsecutivesRecHelper(input, len(input)-1)
}

func MaxSumWithNoTwoConsecutivesRecHelper(input []int, n int) int {
	if n == 0 {
		return input[0]
	}

	if n == 1 {
		return max(input[0], input[1])
	}

	return max(MaxSumWithNoTwoConsecutivesRecHelper(input, n-1), MaxSumWithNoTwoConsecutivesRecHelper(input, n-2)+input[n])
}

func MaxSumWithNoTwoConsecutivesDP(input []int) int {
	memo := make([]int, len(input))

	memo[0] = input[0]
	memo[1] = max(input[0], input[1])

	for i := 2; i < len(input); i++ {
		memo[i] = max(memo[i-1], memo[i-2]+input[i])
	}

	return memo[len(input)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := []int{8, 7, 6, 10}
	fmt.Println(MaxSumWithNoTwoConsecutivesRec(a))
	fmt.Println(MaxSumWithNoTwoConsecutivesDP(a))
}
