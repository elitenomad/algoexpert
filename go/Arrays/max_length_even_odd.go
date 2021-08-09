package main

import "fmt"

func isEven(n int) bool {
	return n&1 == 0
}

func isOdd(n int) bool {
	return n&1 != 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MaxLengthEvenOddBrute(input []int) int {
	res := 1

	for i := 0; i < len(input); i++ {
		current := 1
		for j := i + 1; j < len(input); j++ {
			if (isEven(input[j]) && isOdd(input[j-1])) || (isOdd(input[j]) && isEven(input[j-1])) {
				current += 1
			} else {
				break
			}
		}
		res = max(current, res)
	}
	return res
}

// Below solution is based on Kadanes algorithm
func MaxLengthEvenOddEfficient(input []int) int {
	res := 1
	current := 1

	for i := 1; i < len(input); i++ {
		if (isEven(input[i]) && isOdd(input[i-1])) || (isOdd(input[i]) && isEven(input[i-1])) {
			current += 1
			res = max(current, res)
		} else {
			current = 1
		}
	}

	return res
}

func main() {
	a := []int{5, 10, 20, 6, 3, 8}
	fmt.Println(MaxLengthEvenOddBrute(a))
	fmt.Println(MaxLengthEvenOddEfficient(a))
}
