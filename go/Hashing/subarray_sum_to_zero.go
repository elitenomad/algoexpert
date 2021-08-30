package main

import "fmt"

// array +ve and -ve elements
// Verify if subarray has sum 0 return true
// If not return false

func VerifyIfSubArrayWithSumZero(input []int) bool {

	// Find all the possible sub arrays
	// Return true if there is a sum of 0

	// Prepare a prefix sum array out of it
	// If two of the elements are same we can say that there will be
	// a sub array with prefix 0

	return false
}

func main() {
	a := []int{4, 2, 0, 1}
	fmt.Println(VerifyIfSubArrayWithSumZero(a))
}
