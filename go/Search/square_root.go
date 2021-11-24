package main

import "fmt"

func SquareRoot(n int) int { // O(N)
	i := 1

	/*
		4-8 -> square root is 2
		9-16 => square root is 3
		Our goal is to find the floor of a square root
	*/
	for i*i <= n {
		i++
	}
	return i - 1
}

func SquareRoortUsingBS(n int) int {
	low := 0
	high := n
	result := 0

	for low <= high {
		// Find the mid
		mid := (low + high) / 2

		// Find the Square of a mid
		square := mid * mid

		// Verify its equivalence with square
		if square == n {
			return mid
		}

		// If its greater bring high to mid - 1
		// And do all over again
		// If not move low to mid + 1
		// For every iteration store result of low
		// When the loop gets exited you will have answer
		// In the result
		if square > n {
			high = mid - 1
		} else {
			low = mid + 1
			result = mid
		}
	}

	return result
}

func main() {
	a := 10
	fmt.Println(SquareRoot(a))
	fmt.Println(SquareRoot(16))
	fmt.Println(SquareRoot(1024))
}
