package main

import "fmt"

func Pow(x, n int) int {
	result := 1

	for i := 1; i <= n; i++ {
		result = result * x
	}
	return result
}

func PowRecursive(x, n int) int {
	if n == 0 {
		return 1
	}

	temp := PowRecursive(x, n/2)
	temp = temp * temp
	if n%2 == 0 {
		return temp
	}

	return x * temp
}

func PowerIterative(x, n int) int {
	result := 1

	for n > 0 {
		if n%2 != 0 {
			result = result * x
		}

		x = x * x
		n = n / 2
	}

	return result
}

func main() {
	fmt.Println(Pow(2, 10))
	fmt.Println(PowRecursive(2, 10))
	fmt.Println(PowerIterative(2, 10))
}
