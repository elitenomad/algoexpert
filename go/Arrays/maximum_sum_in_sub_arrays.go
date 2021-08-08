package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//Kadanes algorithm
func MaximumSumOfSubArrays(input []int) int {
	total := input[0]
	n := len(input)

	for i := 0; i < n; i++ {
		curr := 0

		for j := i; j < n; j++ {
			curr = curr + input[j]
			total = max(total, curr)
		}
	}
	return total
}

func KadanesAlgorithm(array []int) int {
	maxAtEachIndex := array[0]
	maxSoFar := array[0]

	for i := 1; i < len(array); i++ {
		val := array[i]
		maxAtEachIndex = max(maxAtEachIndex+val, val)
		maxSoFar = max(maxAtEachIndex, maxSoFar)
	}

	return maxSoFar
}

func main() {
	a := []int{1, -2, 3, -1, 2} // result = 1
	fmt.Println(MaximumSumOfSubArrays(a))
	fmt.Println(KadanesAlgorithm(a))
	a = []int{2, 3, -8, 7, -1, 2, 3}
	fmt.Println(MaximumSumOfSubArrays(a))
	fmt.Println(KadanesAlgorithm(a))
	a = []int{5, 8, 3}
	fmt.Println(MaximumSumOfSubArrays(a))
	fmt.Println(KadanesAlgorithm(a))
}
