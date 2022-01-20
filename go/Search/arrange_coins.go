package main

import "fmt"

func ArrangeCoins(n int) int {
	result := 0
	ArrangeCoinsHelper(n, 1, &result)
	return result
}

func ArrangeCoinsHelper(n int, steps int, result *int) { // O(N)
	if n <= 0 {
		return
	}

	steps += 1
	*result += 1
	ArrangeCoinsHelper(n-steps, steps, result)
}

func ArrangeCoinsBS(n int) int {
	l := 0
	h := n

	for l <= h {
		m := (l + h) / 2

		sum := m * (m + 1) / 2

		if sum == n {
			return m
		}

		if sum > n {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return h
}

func main() {
	n := 10
	fmt.Println(ArrangeCoins(n))
	fmt.Println(ArrangeCoinsBS(n))
}
