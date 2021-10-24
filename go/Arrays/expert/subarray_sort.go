package main

import (
	"fmt"
	"math"
)

/*
	Write a function that takes in an array of atleast 2 integers and that returns an array of the starting and
	Ending indices of the smallest subarray in the input array that needs to be sorted in place in order for the
	entire input array to be sorted (in ascending order)

	if the input is not present return []{-1, -1}
*/

func SubArraySort(input []int) []int {
	// First goal is to find the smallest and largest numbers of
	// the given array
	smallest, largest := math.MaxInt32, math.MinInt32

	for i, num := range input {
		if IsElementsOutOfPlace(input, num, i) {
			smallest = min(smallest, num)
			largest = max(largest, num)
		}
	}

	// Use case when you dont have any smallest or largest
	if smallest == math.MaxInt32 || largest == math.MinInt32 {
		return []int{-1, -1}
	}

	left := 0
	for smallest >= input[left] {
		left += 1
	}

	right := len(input) - 1
	for largest <= input[right] {
		right -= 1
	}

	return []int{left, right}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func IsElementsOutOfPlace(array []int, num, i int) bool {
	if i == 0 {
		return num > array[i+1]
	}

	if i == len(array)-1 {
		return num < array[i-1]
	}

	return num > array[i+1] || num < array[i-1]
}

func main() {
	a := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	fmt.Println(SubArraySort(a))
}
