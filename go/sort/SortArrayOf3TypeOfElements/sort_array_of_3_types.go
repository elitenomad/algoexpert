package main

import "fmt"

/*
	Elements: 0, 1 and 2
	Time complexity : O(n)
	Space complexity : O(n)
*/
func SortNaiveSolution(input []int) []int {
	result := make([]int, len(input))
	i := 0
	for j := 0; j < len(input); j++ {
		if input[j] == 0 {
			result[i] = input[j]
			i++
		}
	}

	for j := 0; j < len(input); j++ {
		if input[j] == 1 {
			result[i] = input[j]
			i++
		}
	}

	for j := 0; j < len(input); j++ {
		if input[j] == 2 {
			result[i] = input[j]
			i++
		}
	}

	return result
}

/*
	Dutch National Flag algorithm
*/
func Sort(input []int) []int {

	return nil
}

func main() {
	a := []int{0, 1, 1, 1, 2, 0, 2, 1, 2, 1, 0}
	fmt.Println(SortNaiveSolution(a))
}
