package main

import "fmt"

// array +ve and -ve elements
// Verify if subarray has sum 0 return true
// If not return false

// Find all the possible sub arrays
// Return true if there is a sum of 0
// Above method requires O(n^2)

// Prepare a prefix sum array out of it
// If two of the elements are same we can say that there will be
// a sub array with prefix 0

func VerifyIfSubArrayWithSumZero(input []int) bool {
	prefixSum := 0
	h := map[int]bool{}

	for i := 0; i < len(input); i++ {
		prefixSum += input[i]

		if input[i] == 0 || h[prefixSum] || prefixSum == 0 {
			return true
		}
		h[prefixSum] = true
	}

	return false
}

func main() {
	a := []int{0, 2, 4, 1}
	fmt.Println(VerifyIfSubArrayWithSumZero(a))
	a = []int{-3, 2, 1}
	fmt.Println(VerifyIfSubArrayWithSumZero(a))
}
