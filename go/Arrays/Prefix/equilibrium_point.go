package main

import "fmt"

func EquilibriumPoint(input []int) bool {
	// Find prefix and suffix arrays
	prefix := []int{}
	suffix := make([]int, len(input))

	sum_prefix := 0
	sum_suffix := 0
	for i, j := 0, len(input)-1; i < len(input) || j >= 0; i, j = i+1, j-1 {
		sum_prefix += input[i]
		sum_suffix += input[j]
		prefix = append(prefix, sum_prefix)
		suffix[j] = sum_suffix
	}

	// fmt.Println(prefix)
	// fmt.Println(suffix)
	for i := 1; i < len(input); i++ {
		if prefix[i-1] == suffix[i+1] {
			return true
		}
	}

	return false
}

func EquilibriumPointEff(input []int) bool {
	total := 0
	for i := 0; i < len(input); i++ {
		total += input[i]
	}

	left_sum := 0
	for i := 0; i < len(input); i++ {
		if left_sum == total-input[i] {
			return true
		}
		left_sum += input[i]
		total -= input[i]
	}

	return false
}

func main() {
	a := []int{2, 8, 3, 9, 6, 5, 4}
	fmt.Println(EquilibriumPointEff(a))
	a = []int{3, 4, 8, -9, 20, 6}
	fmt.Println(EquilibriumPointEff(a))
	a = []int{4, 2, 2}
	fmt.Println(EquilibriumPointEff(a))
}
