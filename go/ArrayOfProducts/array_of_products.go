package main

import "fmt"

// O(n^2) - Best , Worst , Average
// O(n)
func brute(input []int) []int {
	var result []int

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i != j {
				result = append(result, input[i]*input[j])
			}
		}
	}

	return result
}

// Time Complexity - O(n)
// Space Complexity - O(n)
func efficient(input []int) []int {
	var result []int
	for i := 0; i < len(input); i++ {
		result = append(result, 1)
	}
	// Collect the Producct of elements left of the array
	left := 1
	for i := 0; i < len(input); i++ {
		result[i] = left
		left *= input[i]
	}

	right := 1
	for j := len(input) - 1; j >= 0; j-- {
		result[j] *= right
		right *= input[j]
	}

	return result
}

func main() {
	i := []int{5, 1, 4, 2}
	fmt.Println(efficient(i))
}
