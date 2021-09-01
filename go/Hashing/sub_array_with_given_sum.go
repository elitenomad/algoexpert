package main

import "fmt"

func SubArrayWithGivenSum(input []int, sum int) bool {
	prefixSum := 0
	h := map[int]bool{}

	for i := 0; i < len(input); i++ {
		prefixSum += input[i]
		if prefixSum == sum || h[prefixSum-sum] {
			return true
		}

		h[prefixSum] = true
	}

	return false
}

func main() {
	a := []int{5, 8, 6, 13, 3, -1}
	fmt.Println(SubArrayWithGivenSum(a, 5))
}
