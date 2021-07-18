package main

import "fmt"

func sum_of_n_natural_numbers(n int) int {
	if n <= 1 {
		return n
	}

	return n + sum_of_n_natural_numbers(n-1)
}

func main() {
	fmt.Println(sum_of_n_natural_numbers(10))
}
