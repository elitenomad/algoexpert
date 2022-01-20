package main

import "fmt"

func hasDuplicateValue(input []int) bool {
	h := map[int]bool{}

	for _, val := range input {
		if h[val] {
			return true
		} else {
			h[val] = true
		}
	}

	return false
}

// Given unsorted array of integers (NB. it can be array of negative, 0, positive values)
// return a list of duplicate values ?
// Input: [122, 32, 1, 0, 32, -22, -32, 44, -22, 0, -22]
// Expected output: [32, -22, 0]
// -100000 < arr[i] < 100000
func ReturnDuplicateValues(input []int) []int {
	result := []int{}

	return result
}

func main() {
	fmt.Println("Duplicate Linear function implemented:")
	in := []int{1, 1, 2, 3, 4, 5, 6}
	out := hasDuplicateValue(in)
	fmt.Println("In has duplicate value: ", out)
	in = []int{1, 2, 3, 4, 5, 6}
	out = hasDuplicateValue(in)
	fmt.Println("In has duplicate value: ", out)
}
