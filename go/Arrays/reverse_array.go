package main

import "fmt"

func Reverse(input []int) []int {
	// i = 0, j = 4
	for i, j := 0, len(input)-1; i < len(input)/2 || j >= len(input)/2; i, j = i+1, j-1 {
		// fmt.Println("[", i, j, "]")
		input[i], input[j] = input[j], input[i]
	}

	/*
		low := 0
		high := len(input) - 1
		for low < high {
			temp = input[low]
			input[low] = input[high]
			input[high] = temp
			low++
			high--
		}
	*/

	return input
} // O(n/2)

func main() {
	a := []int{4, 3, 2, 10, 9} // 9, 10, 2, 3, 4
	fmt.Println(Reverse(a))
	a = []int{4, 3, 2, 10} // 10, 2,3,4
	fmt.Println(a)
}
