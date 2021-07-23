package main

import "fmt"

func Partition(input []int, start, end int) int {
	i := start - 1

	// Find the length of an array
	// len := len(input) - 1
	// To start with we will consider our pivot as the last element in the array
	pivot := end

	// Loop through the array
	for j := 0; j < end; j++ {
		if input[j] < input[pivot] { // if an element < pivot increment i and swap values
			i++
			input[j], input[i] = input[i], input[j]
		}
	}
	input[i+1], input[end] = input[end], input[i+1]
	return i + 1
}

func QuickSortLomuto(input []int, start, end int) []int {
	if start < end {
		p := Partition(input, start, end)
		QuickSortLomuto(input, start, p-1)
		QuickSortLomuto(input, p+1, end)
	}

	return input
}

func main() {
	i := []int{3, 1, 2, 4, 5, 6}
	fmt.Println(QuickSortLomuto(i, 0, 5))
}
