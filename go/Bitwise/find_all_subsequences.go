package main

import "fmt"

func powOf2(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 2
	}
	return result
}

func FindAllSubsequences(input []int) { // [1,3,5] Not complete yet
	for i := 0; i < (1 << len(input)); i++ {
		// i is the number we draw decsions
		for j := 0; j < len(input); j++ {
			number := (1 << j)

			if i&number != 0 {
				fmt.Printf([]int{input[j]})
			}
		}
	}
}

func main() {
	FindAllSubsequences([]int{1, 3, 5})
}
