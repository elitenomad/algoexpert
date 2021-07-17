package main

import (
	"fmt"
	"math"
)

func CountLogarithmic(n int) int {
	return int(math.Log10(float64(n))) + 1
}

func CountRecursive(n int) int {
	if n == 0 {
		return 0
	}

	return CountRecursive(n/10) + 1
}

func Count(n int) int {
	c := 0

	for n != 0 {
		c += 1
		n = n / 10
	}

	return c
}

func main() {
	n_1 := 123
	fmt.Println(Count(n_1))
	fmt.Println(CountRecursive(n_1))
	fmt.Println(CountLogarithmic(n_1))

	n_2 := 1
	fmt.Println(Count(n_2))
	fmt.Println(CountRecursive(n_2))
	fmt.Println(CountLogarithmic(n_2))

	n_3 := 12343344455
	fmt.Println(Count(n_3))
	fmt.Println(CountRecursive(n_3))
	fmt.Println(CountLogarithmic(n_3))
}
