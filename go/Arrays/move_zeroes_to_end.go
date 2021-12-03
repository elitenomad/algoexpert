package main

import "fmt"

func MoveZeroesToEnd(input []int) []int {
	j := 0

	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			input[i], input[j] = input[j], input[i]
			j += 1
		}
	}

	return input
}

func MoveZeroesToFirst(input []int) []int {
	j := len(input) - 1

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != 0 {
			input[i], input[j] = input[j], input[i]
			j -= 1
		}
	}

	return input
}

func main() {
	a := []int{8, 5, 0, 10, 0, 20}
	fmt.Println(MoveZeroesToEnd(a))
	fmt.Println(MoveZeroesToFirst(a))
	a = []int{0, 0, 0, 5, 10, 30}
	fmt.Println(MoveZeroesToEnd(a))
	fmt.Println(MoveZeroesToFirst(a))
}
