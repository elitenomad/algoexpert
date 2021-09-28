package main

import "fmt"

// T(n) = T(n) + T(n+1) + T(n+3)
func Tribonacci(n int) int { // n <= 13 https://leetcode.com/problems/n-th-tribonacci-number/solution/
	nums := make([]int, n+1)

	if n == 0 {
		nums[0] = 0
		return nums[0]
	}

	nums[1], nums[2] = 1, 1

	for i := 3; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2] + nums[i-3]
	}

	return nums[n]
}

func TribonacciSolutionII(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		} else {
			return n
		}
	}

	x, y, z := 0, 1, 1
	for i := 0; i < n-2; i++ {
		x, y, z = y, z, x+y+z
	}

	return z
}

func main() {
	n := 37
	fmt.Println(Tribonacci(n))
	fmt.Println(TribonacciSolutionII(9))
}
