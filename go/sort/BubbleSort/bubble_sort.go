package main

import (
	"fmt"
)

/*
	Given array of Un-sorted Integers, bubble sort will
	loop through the array and arrange them in ascending order

	For each loop the Largest value goes and sits at the END OF THE ARRAY
*/
func bubble(input []int) []int {
	custom_index := len(input) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < custom_index; i++ {
			if input[i] > input[i+1] {
				temp := input[i+1]
				input[i+1] = input[i]
				input[i] = temp
				sorted = false
			}
		}

		custom_index -= 1
	}

	return input
}

func main() {
	fmt.Println("Bubble Sort Implementation...")
	in := []int{4, 2, 7, 1, 3}
	out := bubble(in)
	fmt.Println("Output is  :", out)
}
