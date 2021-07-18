package main

import "fmt"

func PrintNumbersDec(n int) {
	if n == 1 {
		fmt.Printf("%d", 1)
		return
	}

	fmt.Printf("%d ", n)
	PrintNumbersDec(n - 1)
}

func PrintNumbersInc(n int) { // non tail recursive function
	if n == 0 {
		return
	}

	PrintNumbersInc(n - 1)
	fmt.Printf(" %d", n)
}

func PrintNumbersIncTR(n, k int) {
	if n == 0 {
		return
	}

	fmt.Printf(" %d", k)
	PrintNumbersIncTR(n-1, k+1)
}

func main() {
	PrintNumbersDec(10)
	fmt.Println()
	PrintNumbersInc(10)
	fmt.Println()
	PrintNumbersIncTR(10, 1)
}
