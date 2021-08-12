package main

import "fmt"

func SortedSquareArray(input []int) []int { // O(n) and O(1)
	result := make([]int, len(input))

	sIdx := 0
	eIdx := len(input) - 1

	for i := len(input) - 1; i >= 0; i-- {
		sValue := input[sIdx]
		eValue := input[eIdx]

		if abs(sValue) > abs(eValue) {
			result[i] = sValue * sValue
			sIdx++
		} else {
			result[i] = eValue * eValue
			eIdx--
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func main() {
	a := []int{1, 2, 3, 5, 8, 9}
	fmt.Println(SortedSquareArray(a))
	a = []int{1, 2, 3, 5, 8, 9, 10}
	fmt.Println(SortedSquareArray(a))
	a = []int{-4, -3, -2, -1}
	fmt.Println(SortedSquareArray(a))
	a = []int{-4, 1, 2, 3}
	fmt.Println(SortedSquareArray(a))
}
