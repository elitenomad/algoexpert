package main

import (
	"fmt"
	"sort"
)

func MissingNumber(input []int) int {
	result_1 := 0
	result_2 := 0
	size := len(input)

	for i := 0; i < size; i++ {
		// fmt.Println("Result 1: ", i)
		result_1 = result_1 ^ input[i]
	}

	for i := 1; i <= size+1; i++ {
		// fmt.Println("Result 2: ", i)
		result_2 = result_2 ^ i
	}

	return result_1 ^ result_2
}

func MissingNumberNaive(input []int, n int) int {
	total := make([]int, n)
	result := 0
	for i := 1; i <= n; i++ { // 1,2,3,4
		total[i-1] = i
	}

	sort.Ints(input)
	for i, val := range total {
		if val != input[i] {
			result = val
			break
		}
	}

	return result
}

func MissingNumbers(input []int) []int {
	total := []int{}

	for i, target := 0, input[0]; target <= input[len(input)-1]; target++ { // 3,4,7,8
		if input[i] != target {
			total = append(total, target)
		} else {
			i++
		}
	}

	return total
}

func main() {
	a := []int{1, 4, 3}
	fmt.Println(MissingNumber(a))
	// a = []int{1, 4, 3}
	// fmt.Println(MissingNumberNaive(a, 4))

	// a = []int{1, 5, 3, 2}
	// fmt.Println(MissingNumber(a))
	// a = []int{1, 5, 3, 2}
	// fmt.Println(MissingNumberNaive(a, 5))

	// a = []int{2, 5, 1, 4, 9, 6, 3, 7}
	// fmt.Println(MissingNumber(a))

	a = []int{3, 4, 7, 8}
	fmt.Println(MissingNumbers(a))
}
