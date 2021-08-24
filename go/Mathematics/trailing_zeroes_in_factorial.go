package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func trailingZeroesInFactorial(n int) int { // O(M + N)
	/*
		Find the factorial of a given number
		then do n / 10 until we ran out of 0's
		return the count of all the zeros

		// Complexity: Factorial complexity O(N) + Go through number O(N)
	*/
	q := factorial(n) // O(N) and O(N) Time and Space complexities
	fmt.Println("factorial: ", q)

	count := 0
	r := q % 10

	for r == 0 { // O(LOGM)
		count += 1

		q = q / 10
		r = q % 10
	}

	return count
}

/*
	Goal is to look for no of 2 and 5 as they are
	responsible for adding a 0. Technically no of 2s will
	be lesser than 5's. Special cases when multiple 5s cons
	titute a number.
*/
func trailingZeroesEfficient(n int) int { // O(LOGN), O(1) Time and Space complexities
	p := 5
	noOfZeroes := 0

	for p <= n {
		noOfZeroes += n / p
		p = p * 5
	}
	return noOfZeroes
}

func main() {
	n_1 := 100
	fmt.Println(trailingZeroesEfficient(n_1))
}
