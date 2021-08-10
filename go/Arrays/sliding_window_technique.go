package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Given an array and an integer k
// Find the maximum sum of k consecutive elements
func MaxSumOfkIntegers(input []int, k int) int { // As N reaches K , this will be O(N^2)
	result := 0

	for i := 0; i < len(input)-k; i++ {
		sum := 0

		for j := i; j < i+k; j++ {
			sum += input[j]
		}

		result = max(result, sum)
	}

	return result
}

func MaxSumOfkIntegersEfficient(input []int, k int) int {
	result := 0
	current_max := 0

	for i := 0; i < k; i++ {
		current_max += input[i]
	}

	result = max(result, current_max)

	for i := k; i < len(input); i++ {
		current_max += input[i] - input[i-k]
		result = max(result, current_max)
	}

	return result
}

func main() {
	a := []int{1, 8, 30, -5, 20, 7}
	fmt.Println(MaxSumOfkIntegers(a, 2))
	fmt.Println(MaxSumOfkIntegersEfficient(a, 2))
	fmt.Println(MaxSumOfkIntegers(a, 3))
	fmt.Println(MaxSumOfkIntegersEfficient(a, 3))
}
