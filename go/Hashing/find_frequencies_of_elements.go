package main

import "fmt"

func PrintFrequenciesOfElements(input []int) { // O(N) and O(N) for Time and Space complexities
	h := map[int]int{}

	for _, val := range input {
		h[val] += 1
	}

	for k, v := range h {
		fmt.Println("Key: ", k, " Value : ", v)
	}
}

func main() {
	a := []int{10, 20, 20, 30, 10}
	PrintFrequenciesOfElements(a)
}
