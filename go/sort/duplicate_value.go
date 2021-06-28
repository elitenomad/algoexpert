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

func main() {
	fmt.Println("Duplicate Linear function implemented:")
	in := []int{1, 1, 2, 3, 4, 5, 6}
	out := hasDuplicateValue(in)
	fmt.Println("In has duplicate value: ", out)
	in = []int{1, 2, 3, 4, 5, 6}
	out = hasDuplicateValue(in)
	fmt.Println("In has duplicate value: ", out)
}
