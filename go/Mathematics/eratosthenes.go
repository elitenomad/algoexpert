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

func PrintAllPrimes(n int) int { // Brute force solution
	// Use the Prime(n int) which is written
	// Earlier program
	// Loop until n
	// Verify if the number is primt or not
	// Print the element
	// O(N * Sqrt(N))

	return 0
}

func PrintPrimesUsingEratosthenese(n int) { // Just use fmt.Println to display the numbers
	if n <= 1 {
		return
	}

	var isPrimeSlice = make([]bool, n+1)
	for i := range isPrimeSlice {
		isPrimeSlice[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if PrimeOptimized(i) {
			for j := 2 * i; j <= n; j += i {
				isPrimeSlice[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if isPrimeSlice[i] {
			fmt.Println(i)
		}
	}
}

func main() {
	n := 18
	PrintPrimesUsingEratosthenese(n)
}
