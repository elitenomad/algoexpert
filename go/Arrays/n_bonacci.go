package main

import "fmt"

func nbonacci(n, m int) []int {
	result := []int{}

	// For every n there will be starting n - 1 zeroes
	for j := 0; j < n-1; j++ {
		result = append(result, 0)
	}

	// Add an element 1 to the initial element prepared
	result = append(result, 1)

	// For the rest of the elements until < m , find nbonacci numbers
	for i := n; i < m; i++ {
		sum := 0

		for j := i - 1; j >= i-n; j-- {
			sum += result[j]
		}

		result = append(result, sum)
	}

	return result
}

// N represent what sort of (Fi)Nbonacci
// M Represent number of numbers we should print in sequence.
func nbonacciSlidingWindow(n, m int) []int {
	result := []int{}

	// For every n there will be starting n - 1 zeroes
	for j := 0; j < n-1; j++ {
		result = append(result, 0)
	}

	// Add an element 1 to the initial element prepared
	result = append(result, 1)
	result = append(result, 1)

	// For the rest of the elements until < m , find nbonacci numbers
	for i := n + 1; i < m; i++ { // [0 0 1 1 2 4 7 13]
		sum := 2*result[i-1] - result[i-n-1] // 2 * 1 - result[4 - 3- 1]  -> 2
		result = append(result, sum)         // 2 * result[4] - result[5 - 3 - 1] -> 4
		// 2 * result[5] - result[6 - 3 - 1] -> 8 - 1 -> 7
	}

	return result
}

func main() {
	fmt.Println(nbonacci(3, 8))
	fmt.Println(nbonacci(2, 8))
	fmt.Println(nbonacciSlidingWindow(5, 15))
	fmt.Println()
	fmt.Println(nbonacciSlidingWindow(2, 8))
	fmt.Println(nbonacciSlidingWindow(3, 8))
	fmt.Println(nbonacciSlidingWindow(5, 15))
}
