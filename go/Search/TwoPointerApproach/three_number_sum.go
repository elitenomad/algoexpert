package main

import (
	"fmt"
	"sort"
)

func ThreeNumSum(input []int, target int) []int { //O(N^2) and O(1)
	result := []int{}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == target {
					return []int{input[i], input[j], input[k]}
				}
			}
		}
	}

	return result
}

func TwoNumSum(input []int, target int) []int { //O(N) and O(1)
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

func contains(ints []int, val int) bool {
	for _, v := range ints {
		if v == val {
			return true
		}
	}

	return false
}

func ThreeNumSumEfficient(input []int, target int) [][]int { //O(N^2) and O(N)
	result := [][]int{}
	sort.Ints(input)

	for i := 0; i < len(input); i++ {
		difference := target - input[i]
		remains := TwoNumSum(input, difference)

		if len(remains) > 0 && input[i]+remains[0]+remains[1] == target {
			r := []int{input[i], remains[0], remains[1]}
			result = append(result, r)
		}
	}

	return result
}

func ThreeNumberSumExpert(input []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(input)

	for i := 0; i < len(input)-2; i++ {
		left := i + 1
		right := len(input) - 1

		for left < right {
			sum := input[i] + input[left] + input[right]
			if sum == target {
				r := []int{input[i], input[left], input[right]}
				result = append(result, r)
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}

	}

	return result
}

func main() {
	// a := []int{3, 5, 9, 2, 8, 10, 11}
	// fmt.Println(ThreeNumSum(a, 27))
	a := []int{12, 3, 1, 2, -6, 5, -8, 6}
	fmt.Println(ThreeNumSumEfficient(a, 0))
}
