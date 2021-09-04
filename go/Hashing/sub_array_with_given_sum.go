package main

import "fmt"

// Below function will let us know if sub array exists for a given sum
func SubArrayWithGivenSumExists(input []int, sum int) bool {
	prefixSum := 0
	h := map[int]bool{}

	for i := 0; i < len(input); i++ {
		prefixSum += input[i]
		if prefixSum == sum || h[prefixSum-sum] {
			return true
		}

		h[prefixSum] = true
	}

	return false
}

// Find the subarray with max length given sum
func SubArrayWithGivenSum(input []int, sum int) int {
	// Look into longest sub array with a given sum
	return -1
}

func main() {
	a := []int{5, 8, 6, 13, 3, -1}
	fmt.Println(SubArrayWithGivenSumExists(a, 5))
}
