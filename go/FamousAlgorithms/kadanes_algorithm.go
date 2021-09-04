package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// In efficient way to find the max sub array sum
func InEfficient(input []int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		sum := 0
		for j := i; j < len(input); j++ {
			sum += input[j]
			result = max(result, sum)
		}
	}

	return result
}

func KadanesAlgorithm(input []int) int {
	maxAtEachIndex := input[0]
	maxSoFar := input[0]

	for i := 1; i < len(input); i++ {
		maxAtEachIndex = max(maxAtEachIndex+input[i], input[i])
		maxSoFar = max(maxSoFar, maxAtEachIndex)
	}

	return maxSoFar
}

func main() {
	a := []int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}
	fmt.Println(KadanesAlgorithm(a))
	fmt.Println(InEfficient(a))
}
