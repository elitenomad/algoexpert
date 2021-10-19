package main

import "fmt"

func SearchSortedRowsAndColumns(input [][]int, k int) bool {
	// Adding two conditions
	// Verify if the value to be searched is less than the top left - Return false if its not
	// Verify if the value to be searched is greater than the bottom right - Return false if its not
	rows := len(input) - 1
	columns := 0

	for rows >= columns {
		if input[rows][columns] > k {
			rows -= 1
		} else if input[rows][columns] < k {
			columns += 1
		} else {
			fmt.Println(input[rows][columns])
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(SearchSortedRowsAndColumns([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}, 18))
}
