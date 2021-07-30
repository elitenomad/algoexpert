package main

import "fmt"

/*
	Requires more iterations
*/
func PowerOf2(n int) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n%2 != 0 { // n & 1 == 1
			return false
		}

		n = n / 2 // n >> 1
	}

	return true
}

/*
	Brian kannigams algorithm
*/
func PowerOf2Brain(n int) bool {
	if n == 0 {
		return false
	}

	n = (n & (n - 1))
	if n != 0 {
		return false
	}

	return true
}

func main() {
	a := 13
	b := 16
	fmt.Println(PowerOf2(a), PowerOf2(b))
	fmt.Println(PowerOf2Brain(a), PowerOf2Brain(b))
}
