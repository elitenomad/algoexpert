package main

import "fmt"

func PrintBoundaries(input [][]int) {
	rows := len(input)
	columns := len(input[0])

	r := 0
	c := 0

	if rows == 1 {
		for i := 0; i < columns; i++ {
			fmt.Printf("%d ", input[0][i])
		}
		fmt.Println()
		return
	} else if columns == 1 {
		for i := 0; i < rows; i++ {
			fmt.Printf("%d ", input[i][0])
		}
		fmt.Println()
		return
	}

	for true {

		for c < columns-1 {
			fmt.Printf("%d ", input[r][c])
			c += 1
		}

		for r < rows-1 {
			fmt.Printf("%d ", input[r][c])
			r += 1
		}

		for c > 0 {
			fmt.Printf("%d ", input[r][c])
			c -= 1
		}

		for r > 0 {
			fmt.Printf("%d ", input[r][c])
			r -= 1
		}

		if r == 0 && c == 0 {
			break
		}
	}
	fmt.Println()
}

func main() {
	PrintBoundaries([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})

	PrintBoundaries([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})

	PrintBoundaries([][]int{
		{1, 2, 3, 4},
	})

	PrintBoundaries([][]int{
		{1},
		{2},
		{3},
	})

	PrintBoundaries([][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	})
}
