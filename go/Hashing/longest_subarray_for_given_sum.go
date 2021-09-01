package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func LongestSubArrayForGivenSum(input []int, n int) int {
	result := 0
	prefixSum := 0
	h := map[int]int{}

	for i := 0; i < len(input); i++ {
		prefixSum += input[i]

		if prefixSum == n {
			result = i + 1
		}

		if _, found := h[prefixSum-n]; found {
			result = max(result, i-h[prefixSum-n])
		}

		h[prefixSum] = i
	}

	return result
}

func main() {
	a := []int{5, 8, -4, -4, 9, -2, 2}
	fmt.Println(LongestSubArrayForGivenSum(a, 0))
	a = []int{8, 3, 7}
	fmt.Println(LongestSubArrayForGivenSum(a, 15))
	a = []int{3, 1, 0, 1, 8, 2, 3, 6}
	fmt.Println("***")
	fmt.Println(LongestSubArrayForGivenSum(a, 5))
}
