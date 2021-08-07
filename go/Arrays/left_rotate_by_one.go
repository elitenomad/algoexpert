package main

import "fmt"

func LeftRotateByOne(input []int) []int {
	first := input[0]

	for i := 1; i < len(input); i++ { // O(n)
		input[i-1] = input[i]
	}

	input[len(input)-1] = first

	return input
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(LeftRotateByOne(a))
	a = LeftRotateByOne(a)
	fmt.Println(a)
}
