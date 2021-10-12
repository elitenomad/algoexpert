package main

import "fmt"

func LeftRotateByDSpaces(input []int, d int) []int { // O(N * D)

	for d > 0 {
		first := input[0]

		for i := 1; i < len(input); i++ { // O(n)
			input[i-1] = input[i]
		}

		input[len(input)-1] = first
		d--
	}

	return input
}

func LeftRotateByDSpacesII(input []int, d int) []int { // O(n) and O(d) Time and space
	result := []int{}

	result = input[0:d]
	input = input[d:] // Instead of this we can do below as well
	// for i := d; i < len(input); i++ {
	// 	input[i-d] = input[i]
	// }

	for _, val := range result {
		input = append(input, val)
	}

	return input
}

func reverse(input []int, l, h int) []int {
	for l < h {
		input[l], input[h] = input[h], input[l]
		l++
		h--
	}

	return input
}

func LeftRotateByDSpacesEfficient(input []int, d int) []int {
	input = reverse(input, 0, d-1)
	input = reverse(input, d, len(input)-1)
	input = reverse(input, 0, len(input)-1)

	return input
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(LeftRotateByDSpaces(a, 3))
	a = []int{1, 2, 3, 4, 5}
	fmt.Println(LeftRotateByDSpacesII(a, 3))
	a = []int{1, 2, 3, 4, 5}
	fmt.Println(LeftRotateByDSpacesEfficient(a, 2))
}
