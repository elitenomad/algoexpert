package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func LongestSubArrayInEfficient(input []int) int {
	result := 0

	for i := 0; i < len(input); i++ {
		c1 := 0
		c0 := 0

		for j := i; j < len(input); j++ {
			if input[j] == 0 {
				c0 += 1
			}

			if input[j] == 1 {
				c1 += 1
			}

			if c1 == c0 {
				result = max(result, c0+c1)
			}
		}
	}

	return result
}

func LongestSubArrayForGivenSum(input []int, sum int) int {
	result := 0
	prefixSum := 0
	end_index := -1
	h := map[int]int{}

	for i := 0; i < len(input); i++ {
		prefixSum += input[i]

		if prefixSum == sum {
			result = i + 1
		}

		if _, found := h[prefixSum-sum]; found {
			fmt.Println(":s ::", i-h[prefixSum-sum])
			result = max(result, i-h[prefixSum-sum])
		} else {
			h[prefixSum] = i
		}
	}

	return result
}

func LongestSubArrayOfEqual0sAnd1s(input []int) int {
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			input[i] = -1
		}
	}

	return LongestSubArrayForGivenSum(input, 0)
}

func main() {
	a := []int{1, 1, 1, 0, 1, 0, 1, 1, 1}
	fmt.Println(LongestSubArrayInEfficient(a))
	fmt.Println(LongestSubArrayOfEqual0sAnd1s(a))
}
