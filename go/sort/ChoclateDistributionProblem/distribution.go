package main

import (
	"fmt"
	"sort"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Distribution(input []int, n, m int) int {
	if m > n {
		return -1
	}

	// Sort an array
	sort.Sort(sort.IntSlice(input))
	fmt.Println(input)
	min := input[m-1] - input[0]
	for i := 1; (i + m - 1) < n; i++ {
		min = Min(min, input[i+m-1]-input[i])
	}

	return min
}

func main() {
	a := []int{7, 3, 2, 4, 9, 12, 56}
	m := 3
	fmt.Println(Distribution(a, 7, m))
}
