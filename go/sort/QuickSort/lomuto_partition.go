package main

import "fmt"

func LomutoPartition(input []int) int {
	i := -1

	// Find the length of an array
	len := len(input) - 1
	// To start with we will consider our pivot as the last element in the array
	pivot := input[len]

	// Loop through the array
	for j := 0; j <= len; j++ {
		if input[j] < pivot { // if an element < pivot increment i and swap values
			i++
			input[j], input[i] = input[i], input[j]
		}
	}
	input[i+1], input[len] = input[len], input[i+1]
	return i + 1
}

func main() {
	a := []int{3, 8, 6, 12, 10, 7}
	fmt.Println(LomutoPartition(a))
	a = []int{70, 60, 80, 40, 30}
	fmt.Println(LomutoPartition(a))
	a = []int{30, 40, 20, 50, 80}
	fmt.Println(LomutoPartition(a))
}
