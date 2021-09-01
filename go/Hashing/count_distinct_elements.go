package main

import "fmt"

func CountDistinctElements(input []int) int { // O(N) and O(N) for both time and space
	h := map[int]int{}

	for _, val := range input {
		h[val] += 1
	}

	return len(h)
}

func main() {
	a := []int{15, 12, 13, 12, 13, 13, 18}
	fmt.Println(CountDistinctElements(a))
	a = []int{10, 10, 10}
	fmt.Println(CountDistinctElements(a))
}
