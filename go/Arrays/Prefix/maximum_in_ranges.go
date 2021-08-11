package main

import "fmt"

// Given 0 <= L[i], R[i] <=1000
// We assume l and r are of same length
func MaximumAppearancesInRanges(l, r []int) int {
	// Initialize and array with 1000
	input := make([]int, 1000)
	for i := 0; i < len(l); i++ {
		input[l[i]]++
		input[r[i]+1]--
	}

	fmt.Println(input[:25])
	max := input[0]
	res := 0

	for i := 1; i < 1000; i++ {
		input[i] += input[i-1]
		if max < input[i] {
			max = input[i]
			res = i
		}
	}
	fmt.Println(input[:25])
	return res
}

func main() {
	l := []int{1, 2, 5, 15}
	r := []int{5, 8, 7, 18}
	fmt.Println(MaximumAppearancesInRanges(l, r)) // Output: 5
	l = []int{1, 2, 3}
	r = []int{4, 5, 7}
	fmt.Println(MaximumAppearancesInRanges(l, r)) // Output: 3
}
