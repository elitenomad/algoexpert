package main

import (
	"fmt"
	"sort"
)

func Method1(input []int, k int) {
	sort.Ints(input) // O(NLogN)
	current := 1
	fmt.Println(input)
	// Look for elements with more than n/k occurences and print them
	for i := 0; i < len(input)-1; i++ { //O(N)
		if input[i] == input[i+1] {
			current += 1
		} else {
			if current > (len(input) / k) {
				fmt.Println(input[i])
			}
			current = 1
		}
	}
}

func Method2(input []int, k int) {
	h := map[int]int{}                // Auxillary space used is O(N)
	for i := 0; i < len(input); i++ { //O(N)
		h[input[i]] += 1
	}

	for key, v := range h { // O(N)
		if v > (len(input) / k) {
			fmt.Println(key)
		}
	}
}

func Method3(input []int, k int) {
	// Think about O(NK) solution
}

func main() {
	a := []int{10, 10, 20, 30, 20, 10, 10}
	// Method1(a, 2)
	Method2(a, 2)
}
