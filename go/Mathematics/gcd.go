package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Approach - 1
func gcd(a, b int) int { // O(MIN(A, B))
	mn := min(a, b)

	for mn != 0 {
		if a%mn == 0 && b%mn == 0 {
			break
		}

		mn -= 1
	}

	return mn
}

// Approach - 2 = Eucledean algorithm
func gcd_2(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}

	return a
}

//Approach - 3 Optimized version of EA
func gcd_3(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd_3(b, a%b)
}

func main() {
	a, b := 10, 15

	fmt.Println(gcd(a, b))
	fmt.Println(gcd_2(a, b))
	fmt.Println(gcd_3(a, b))
}
