package main

import "fmt"

/*
	input : Array from which element will be deleted
	n : n is the number to be deleted

	Implementation keeps last value as its a dynamic array
*/
func DeleteAnElemFromArray(input []int, n int) []int {
	var j int

	j = -1

	for i, val := range input {
		if val == n {
			j = i
		}
	}

	if j == -1 {
		return input
	}

	for i := j; i < len(input)-1; i++ {
		input[i] = input[i+1]
	}

	return input
}

func main() {
	a := []int{2, 3, 4, 6, 7, 10}
	fmt.Println(DeleteAnElemFromArray(a, 4))
}
