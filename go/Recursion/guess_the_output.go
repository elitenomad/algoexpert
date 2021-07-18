package main

import "fmt"

func GuessTheOutput(n int) {
	if n == 0 {
		return
	}

	fmt.Println("Before N is :", n)
	GuessTheOutput(n - 1)
	fmt.Println("After N is :", n)
}

func GuessTheOutput2(n int) {
	if n <= 0 {
		return
	}

	GuessTheOutput2(n - 1)
	fmt.Println("N is :", n)
	GuessTheOutput2(n - 1)
}

func main() {
	GuessTheOutput2(3)
}
