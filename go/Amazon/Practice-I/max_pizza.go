package main

import (
	"fmt"
	"math"
)

// Cannot eat pizza from adjacent bowls
// Have to find the max number of pizzas you can eat following above rule
func MaxPizzas(input []int, n int) int {
	dp := make([]int, len(input))
	if n == 0 || len(input) <= 0 {
		return 0
	}

	dp[0] = input[0]                // If you have only one element
	dp[1] = max(input[0], input[1]) // If you have two elements you pick max of both

	for i := 2; i < len(input); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+input[i])
	}

	result := math.MinInt32
	for i := 0; i < len(dp); i++ {
		if dp[i] > result {
			result = dp[i]
		}
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
	a := []int{1, 2, 3, 4, 5}
	n := 5
	fmt.Println(MaxPizzas(a, n)) // 9
	a = []int{5, 3, 4, 11, 2}
	n = 5
	fmt.Println(MaxPizzas(a, n)) // 16
}
