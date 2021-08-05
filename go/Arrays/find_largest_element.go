package main

import "fmt"

func LargestElementOfArray(input []int) int {
	result := -1

	for i, val := range input {
		if val > result {
			result = i
		}
	}

	return result
}

func main() {
	a := []int{5, 7, 1, 3, 10}
	fmt.Println(LargestElementOfArray(a))
}
