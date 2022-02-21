package main

import "fmt"

// Goal is to find the number of subsets sum to the number provided
func SubsetSumRec(input []int, sum int, n int) int {
	if n == 0 {
		if sum == 0 {
			return 1
		} else {
			return 0
		}
	}

	return SubsetSumRec(input, sum, n-1) + SubsetSumRec(input, sum-input[n-1], n-1)
}

func main() {
	a := []int{10, 20, 15}
	fmt.Println(SubsetSumRec(a, 25, len(a)))
	fmt.Println(SubsetSumRec(a, 45, len(a)))
}
