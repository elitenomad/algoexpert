package main

import "fmt"

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

func main() {
	a := []int{1, 4, 3}
	fmt.Println(MissingNumber(a))

	a = []int{1, 5, 3, 2}
	fmt.Println(MissingNumber(a))
}
