package main

import "fmt"

// Given set of coins we need to find
// how many ways we can reach to targetSum
func CoinChange(input []int, target int) int {
	return CoinChangeHelper(input, len(input), target)
}

func CoinChangeHelper(input []int, size, target int) int {
	if target <= 0 {
		return 1
	}

	if size == 0 {
		return 0
	}

	// Scenario 1 We recursively call an array
	// with size - 1 and target
	res := CoinChangeHelper(input, size-1, target)
	// fmt.Println(res)
	if input[size-1] <= target {
		res += CoinChangeHelper(input, size, target-input[size-1])
	}

	return res
}

func main() {
	a := []int{1, 2, 3}
	target := 4
	fmt.Println(CoinChange(a, target))
}
