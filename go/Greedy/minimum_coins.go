package main

import (
	"fmt"
	"sort"
)

// How can you account for the given `amount` with
// Minmum amount of coins from input array
func MinimumCoins(input []int, amount int) int {
	result := 0
	sort.Sort(sort.Reverse(sort.IntSlice(input)))

	fmt.Println(input)
	for i := 0; i < len(input); i++ {
		if input[i] <= amount {
			result += (amount / input[i])
			amount = amount % input[i]
		}

		if amount <= 0 {
			break
		}
	}

	return result
}
func main() {
	a := []int{10, 5, 1, 2}
	fmt.Println(MinimumCoins(a, 57))
}
