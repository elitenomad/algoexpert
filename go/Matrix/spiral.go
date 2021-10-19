package main

import "fmt"

func SpiralMatrix(input [][]int) {
	rows := len(input)
	columns := len(input[0])

	// Initialize 4 variables
	top := 0
	bottom := rows - 1
	left := 0
	right := columns - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			fmt.Println(input[top][i])
		}

		top += 1

		for i := top; i <= bottom; i++ {
			fmt.Println(input[i][bottom])
		}

		right -= 1

		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Println(input[bottom][i])
			}
			bottom -= 1
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Println(input[i][left])
			}
			left += 1
		}
	}

}

func main() {
	SpiralMatrix([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
}
