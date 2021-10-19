package main

import "fmt"

func MedianInSortedRowMatrix(input [][]int, k int) int {
	rows := len(input) - 1
	columns := len(input[0]) - 1

	// Find min and max
	min := input[0][0]
	max := input[0][columns]
	for i := 1; i < rows; i++ {
		if input[i][0] < min {
			min = input[i][0]
		}

		if input[i][columns-1] > max {
			max = input[i][columns-1]
		}
	}

	// We know that median of odd number of elemnts
	medPoint := (rows * columns) / 2

	for min < max {
		mid := (min + max) / 2
		midPos := 0

		for i := 0; i < rows; i++ {
			// Find the position using array.binarysearch add 1
		}

		if midPos < medPoint {
			min = mid + 1
		} else {
			max = mid
		}

	}

	return min
}

func main() {
	fmt.Println(MedianInSortedRowMatrix([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}, 18))
}
