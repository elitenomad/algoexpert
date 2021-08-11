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

func MaxSubArrayForSumOfKIntegersExists(input []int, k, sum int) bool {
	current_max := 0

	for i := 0; i < k; i++ {
		current_max += input[i]
	}

	if current_max == sum {
		return true
	}

	for i := k; i < len(input); i++ {
		current_max += input[i] - input[i-k]
		if current_max == sum {
			return true
		}
	}

	return false
}

// Given an array of unsorted Non -ve Integers
// Figure out whether there is a subarray that mounts to
// Given SUM.
// Given an array [8,5,1,3,6,19], Does it have a subarray which
// mounts to the sum of 9 . Ans is true [1,3,5]
func MaxSubArrayExists(input []int, sum int) bool {
	start := 0
	current_sum := input[0]

	for end := 1; end < len(input); end++ {
		// Clear the current_sum if its  > sum given
		for current_sum > sum && start < end-1 {
			current_sum -= input[start]
			start++
		}

		if current_sum == sum {
			return true
		}

		// Add the element to current_sum
		current_sum += input[end]
	}

	fmt.Println("CURRENT_SUM: ", current_sum)
	return current_sum == sum
}

func main() {
	a := []int{1, 8, 30, -5, 20, 7}
	fmt.Println(MaxSumOfkIntegers(a, 2))
	fmt.Println(MaxSumOfkIntegersEfficient(a, 2))
	fmt.Println(MaxSumOfkIntegers(a, 3))
	fmt.Println(MaxSumOfkIntegersEfficient(a, 3))
	fmt.Println(MaxSubArrayForSumOfKIntegersExists(a, 3, 55))
	a = []int{1, 8, 30, 5, 20, 7}
	fmt.Println(MaxSubArrayExists(a, 45))
}
