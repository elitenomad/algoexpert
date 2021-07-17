package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialIterative(n int) int {
	result := 1

	for n > 0 {
		result *= n
		n -= 1
	}

	return result
}

func main() {
	n_1 := 6
	fmt.Println(factorial(n_1))
	n_2 := 6
	fmt.Println(factorialIterative(n_2))
}
