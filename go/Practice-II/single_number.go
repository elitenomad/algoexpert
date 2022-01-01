package main

import "fmt"

/*
	Given an array of integers, every element appears thrice except for one
	which occurs just once. Find that unique element
*/

// Use case 1 : input is empty
// User case 2 : is XOR all the numbers - Repeating elements will turn to 0.
// This is when number is duplicated (and not triplicated)
func UniqueElement(input []int) int {
	if len(input) <= 0 {
		return -1
	}

	result := input[0]

	for i := 1; i < len(input); i++ {
		result = result ^ input[i]
	}

	return result
}

func SingleNumberII(input []int) int {
	// Using hash strategy
	// create a hash map
	// Store and increment the counter of each value
	// Loop through the hash and look for value == 1
	// Return the key of the hash with value 1
	// TC : O(N) and SC : O(N)
	return -1
}

func SingleNumberIII(input []int) int {
	// Create two variables
	// seenOnce and seenTwice
	// Add it to seenOnce if its the first time the value is being seen
	// Remove it from seenOnce if its seen second time, add it to the seenTwice
	// Thirdtime, dont add it to seenOnce as the value is still in seenTwice
	// Remove it from seenTwice

	// Look for leetcode for solution
	return -1
}

func main() {
	a := []int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}
	fmt.Println(UniqueElement(a))
}
