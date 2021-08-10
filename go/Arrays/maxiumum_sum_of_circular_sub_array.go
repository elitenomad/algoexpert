package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MaximumSumOfCircularSubArray(input []int) int {
	result := input[0]

	for i := 0; i < len(input); i++ {
		current_max := input[i]
		current_sum := input[i]

		for j := 1; j < len(input); j++ {
			index := (i + j) % len(input)
			current_sum += input[index]
			current_max = max(current_max, current_sum)
		}

		result = max(result, current_max)
	}

	return result
}

func main() {
	a := []int{5, -2, 3, 4}
	fmt.Println(MaximumSumOfCircularSubArray(a))
}
