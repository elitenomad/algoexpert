package main

import "fmt"

func TwoNumSum(input []int, target int) []int { //O(N^2) and O(1)
	result := []int{}

	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input); j++ {
			if input[i]+input[j] == target {
				return []int{input[i], input[j]}
			}
		}
	}

	return result
}

func TwoNumSumEfficient(input []int, target int) []int { //O(N) and O(1)
	h := make(map[int]bool)

	for i := 0; i < len(input); i++ {
		difference := target - input[i]

		if _, found := h[difference]; found {
			return []int{difference, input[i]}
		}

		h[input[i]] = true
	}

	return []int{}
}

func main() {
	a := []int{3, 5, 9, 2, 8, 10, 11}
	fmt.Println(TwoNumSum(a, 17))
	fmt.Println(TwoNumSumEfficient(a, 17))
}
