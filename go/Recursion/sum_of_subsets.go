package main

import "fmt"

// Given an array , find all the subsets
// Verify if the elements match the sum in each iteration
// If so add 1
func SumOfSubsets(n []int, sum int) int {
	current := []int{}
	var s int

	// Not sure if this is right way to do
	SumOfSubsetsHelper(n, sum, &s, current, 0)

	return s
}

func SumOfSubsetsHelper(n []int, s int, sum *int, current []int, i int) {
	if i == len(n) {
		temp := 0
		// Find the total sum of an arra
		for _, v := range current {
			temp += v
		}
		// Verify the above sum  if its equivalent to the sum passed
		if temp == s {
			*sum += 1
		}
		return
	}

	SumOfSubsetsHelper(n, s, sum, current, i+1)
	current = append(current, n[i])
	SumOfSubsetsHelper(n, s, sum, current, i+1)
}

// Solution without maintaining the subsets
// We just maintain the sum
// O(2 ^ n)
func SumOfSubsets2(n []int, sum int) int {
	return SumOfSubsets2Helper(n, sum, len(n))
}
func SumOfSubsets2Helper(n []int, sum int, s int) int {
	if s == 0 {
		if sum == 0 {
			return 1
		} else {
			return 0
		}
	}

	return SumOfSubsets2Helper(n, sum, s-1) + SumOfSubsets2Helper(n, sum-n[s-1], s-1)
}

func main() {
	n := []int{10, 20, 15}
	sum := 25

	fmt.Println(SumOfSubsets2(n, sum))
}
