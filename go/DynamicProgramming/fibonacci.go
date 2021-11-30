package main

import "fmt"

// TOP Down approach
// Recursion + Caching
func fibonacci(n int) int { // 0 1 1 2 3 5
	m := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		m[i] = -1
	}

	if m[n] == -1 {
		var res int

		if n == 0 {
			res = 0
			return res
		}

		if n == 1 || n == 2 {
			res = 1
			return res
		}

		res = fibonacci(n-1) + fibonacci(n-2)
		m[n] = res
	}

	return m[n]
}

// Bottom Up approach
// Pre calculate everything before
// And lookup when required
func fibonacci_bu(n int) int {
	fib := make([]int, n+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n+1; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	// TOP DONW
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(30))

	// BOTTOM UP
	fmt.Println(fibonacci_bu(4))
	fmt.Println(fibonacci_bu(5))
	fmt.Println(fibonacci_bu(6))
	fmt.Println(fibonacci_bu(30))
}
