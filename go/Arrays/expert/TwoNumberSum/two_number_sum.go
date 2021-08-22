package main

import "fmt"

// When Array is unsorted we can use hash
func TwoNumberSum(array []int, target int) []int { // O(N) and O(N)
	var store = make(map[int]bool)

	for _, elem := range array {
		sub := target - elem

		if _, found := store[sub]; found {
			return []int{sub, elem}
		}

		store[elem] = true
	}

	return []int{}
}

// Two pointer approach for sorted array
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func TwoNumberSum2PointerApproach(input []int, target int) []int {
	i := 0
	j := len(input) - 1
	// {3,5,9,2,8,10,11}, 17
	// {2,3,5,8,9,10,11}, 17

	for i <= j {
		result := abs(input[j] + input[i])

		if result == target {
			return []int{input[i], input[j]}
		}

		if result < target {
			i++
		} else {
			j--
		}
	}

	return []int{}
}

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	givenSum := 10
	fmt.Println(TwoNumberSum(a, givenSum))
	a = []int{2, 3, 5, 8, 9, 10, 11}
	fmt.Println(TwoNumberSum2PointerApproach(a, 17))
	a = []int{3, 5, 9, 2, 8, 10, 11}
	fmt.Println(TwoNumberSum2PointerApproach(a, 10)) // Two pointer approach fails for unsorted array
}
