package main

import (
	"fmt"
)

func NonRepeatingElements(input []int) []int {
	result := []int{}
	h := map[int]int{}

	for i := 0; i < len(input); i++ {
		h[input[i]] += 1
	}

	for k, v := range h {
		if v == 1 {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	a := []int{0, 9, 2, 3, 4, 8, 7}
	fmt.Println(NonRepeatingElements(a))
	a = []int{10, 20, 40, 30, 10}
	fmt.Println(NonRepeatingElements(a))
}
