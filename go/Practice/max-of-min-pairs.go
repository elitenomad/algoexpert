// This is my solution for AlgoDaily problem Max of Min Pairs
// Located at https://algodaily.com/challenges/max-of-min-pairs

package main

import (
	"fmt"
	"sort"
)

func maxOfMinPairs(input []int) int {
	// We will sort our array
	sort.Ints(input)
	result := 0

	for i := 0; i < len(input)-1; i += 2 {
		result += input[i]
	}

	return result
}

func main() {
	// write test cases
	fmt.Println(maxOfMinPairs([]int{1, 3, 2, 6, 5, 5}))
	fmt.Println(maxOfMinPairs([]int{1, 5, 2, 3}))
	fmt.Println(maxOfMinPairs([]int{1, 3, 2, 1, 4, 5}))
}
