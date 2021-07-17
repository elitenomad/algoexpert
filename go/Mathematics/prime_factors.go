package main

import "fmt"

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

	for i := 5; i*i < a; i++ { // instead of i++ , i += 6
		if a%i == 0 || a%(i+2) == 0 {
			return false
		}
	}

	return true
}

func PrimeFactors(a int) {
	for i := 2; i < a; i++ {
		if PrimeOptimized(i) {
			x := i
			for a%x == 0 {
				fmt.Printf("%d \t", i)
				x = x * i
			}
		}
	}

	return
}

func PrimeFactorsOptimized(a int) {
	if a <= 1 {
		return
	}

	for a%2 == 0 {
		fmt.Printf("%d \t", 2)
		a = a / 2
	}

	for a%3 == 0 {
		fmt.Printf("%d \t", 3)
		a = a / 3
	}

	for i := 5; i*i < a; i = i + 6 {
		if PrimeOptimized(i) {
			for a%i == 0 {
				fmt.Printf("%d \t", i)
				a = a / i
			}

			for a/(i+2) == 0 {
				fmt.Printf("%d \t", i+2)
				a = a / (i + 2)
			}
		}
	}

	if a > 3 {
		fmt.Printf("%d", a)
	}

	return
}

func main() {
	_, b := 12, 84

	PrimeFactorsOptimized(b)
}
