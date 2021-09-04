package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Assumption is length of a and b are same
func LongestCommonSubArray(a, b []int) int {
	result := 0

	for i := 0; i < len(a); i++ { // Loop through array
		s_a := 0 //Initialize two sum varaibles
		s_b := 0
		for j := i; j < len(b); j++ {
			s_a += a[j] // Collect sum of subarrays from A
			s_b += b[j] // Collect sum of subarrays from B

			if s_a == s_b { //If its equal collect the result
				result = max(result, j-i+1) // Replace max traversal at each iteration
			}
		}
	}

	return result
}

func LongestSubArrayForGivenSum(input []int, sum int) int {
	result := 0
	prefixSum := 0
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

func LongestCommonSubArrayEff(a, b []int) int {
	// difference of two arrasy
	c := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		c[i] = a[i] - b[i]
	}

	return LongestSubArrayForGivenSum(c, 0)
}

func main() {
	a := []int{0, 1, 0, 0, 0, 0}
	b := []int{1, 0, 1, 0, 0, 1}
	fmt.Println(LongestCommonSubArray(a, b))
	fmt.Println(LongestCommonSubArrayEff(a, b))
}
