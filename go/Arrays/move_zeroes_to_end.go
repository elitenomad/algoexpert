package main

import "fmt"

func MoveZeroesToEnd(input []int) []int {
	size := len(input) - 1

	for i := 0; i < len(input)/2; i++ {
		if input[i] == 0 {
			input[i], input[size] = input[size], input[i]
			// fmt.Println(input)
			size = size - 1
		}
	}

	return input
}

func main() {
	a := []int{8, 5, 0, 10, 0, 20}
	fmt.Println(MoveZeroesToEnd(a))
	a = []int{0, 0, 0, 5, 10, 30}
	fmt.Println(MoveZeroesToEnd(a))
}
