package main

import "fmt"

func OddOccuringNumberNaive(input []int) int {
	h := map[int]int{}
	result := -1
	// Loop through an array
	for i := 0; i < len(input); i++ {
		if _, ok := h[input[i]]; ok {
			h[input[i]] = h[input[i]] + 1
		} else {
			h[input[i]] = 1
		}
	}

	for key, val := range h {
		if val%2 != 0 {
			result = key
		}
	}

	return result
}

func OddOccuringNumber(input []int) int {
	result := 0
	for i := 0; i < len(input); i++ {
		result = result ^ input[i]
	}

	return result
}

func main() {
	a := []int{8, 7, 7, 8, 8}
	fmt.Println(OddOccuringNumberNaive(a))
	fmt.Println(OddOccuringNumber(a))
}
