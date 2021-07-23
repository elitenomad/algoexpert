package main

import "fmt"

func Partition(input []int, first, last int) int {
	i := first - 1
	j := last + 1
	pivot := input[first]

	for {
		for ok := true; ok; ok = (input[i] < pivot) {
			i++
		}

		for ok := true; ok; ok = (input[j] > pivot) {
			j--
		}

		if i >= j {
			return j
		}
		input[i], input[j] = input[j], input[i]
	}
}

func QuickSortHoares(input []int, start, end int) []int {
	if start < end {
		p := Partition(input, start, end)
		QuickSortHoares(input, start, p)
		QuickSortHoares(input, p+1, end)
	}

	return input
}

func main() {
	i := []int{3, 1, 2, 4, 5, 6}
	fmt.Println(QuickSortHoares(i, 0, 5))
}
