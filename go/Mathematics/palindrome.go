package main

import (
	"fmt"
)

func Reverse(n int) int {
	last := 0

	for n != 0 {
		r := n % 10

		last = last * 10

		last = last + r

		n = n / 10
	}

	return last
}

func IsPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	if n >= 0 && n <= 9 {
		return true
	}

	return Reverse(n) == n
}

func main() {
	n := 123
	fmt.Println(IsPalindrome(n))
	n_1 := 313
	fmt.Println(IsPalindrome(n_1))
}
