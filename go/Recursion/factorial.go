package main

import "fmt"

func factorialNTR(n int) int { // Non Tail Recursive function
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorialNTR(n-1)
}

func factorialTR(n, k int) int { // Tail recursive function - pass k as 1
	if n == 0 || n == 1 {
		return 1
	}

	return factorialTR(n-1, k*n)
}

func main() {
	fmt.Println(factorialTR(5, 1))
}
