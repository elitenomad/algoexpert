package main

import "fmt"

func Prime(a int) bool { // 1 is neither composite not prime
	if a == 2 || a == 3 {
		return true
	}

	if a%2 == 0 || a%3 == 0 {
		return false
	}

	n := 4

	for a != n && n < a {
		if a%n == 0 {
			return false
		}

		n += 1
	}

	return true
}

// Factors of a number always come in pairs
// We should be able to find if a number has a factor
// before n reaches its SqrtOfN
// e.g 65 (1, 65), (13, 5)
func PrimeOptimized(a int) bool { // O(SqrtofN)
	if a == 1 {
		return false
	}

	if a == 2 || a == 3 {
		return true
	}

	if a%2 == 0 || a%3 == 0 {
		return false
	}

	for i := 5; i*i <= a; i++ { // instead of i++ , i += 6
		if a%i == 0 || a%(i+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	a, b, c := 2, 3, 10
	fmt.Println(Prime(a))
	fmt.Println(Prime(b))
	fmt.Println(Prime(c))

	fmt.Println(PrimeOptimized(a))
	fmt.Println(PrimeOptimized(b))
	fmt.Println(PrimeOptimized(25))
}
