package main

import (
	"fmt"
	"math"
	"sort"
)

func MinDifference(input []int) float64 {
	// Sort an array
	sort.Ints(input)

	// Large +ve value
	min := math.Inf(+10)

	for i := 1; i < len(input); i++ {
		diff := (input[i] - input[i-1])
		min = math.Min(float64(min), float64(diff))
	}

	return min
}

func main() {
	a := []int{2, 3, 1, 4, 6}
	fmt.Println(MinDifference(a))
	a = []int{1, 8, 12, 5, 18}
	fmt.Println(MinDifference(a))
	a = []int{8, 15}
	fmt.Println(MinDifference(a))
	a = []int{8, -1, 0, 3}
	fmt.Println(MinDifference(a))
	a = []int{10}
	fmt.Println(MinDifference(a))
}
