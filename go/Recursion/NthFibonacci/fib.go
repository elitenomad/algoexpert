package main

import "fmt"

// 0,1,1,2,3,5,8
func fib(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 || n == 3 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(5))
}
