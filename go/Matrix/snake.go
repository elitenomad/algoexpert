package main

import "fmt"

func PrintInSnakePattern(input [][]int) {
	rows := len(input)
	columns := len(input[0])

	i := 0
	snake := 1 // 1 forward -1 represents backward

	for i < rows {
		if snake == 1 {
			j := 0
			for j = 0; j < columns; j++ {
				fmt.Printf("%d \t", input[i][j])
			}

			if j == columns {
				snake = -1
			}
		} else if snake == -1 {
			j := columns - 1
			for j = columns - 1; j >= 0; j-- {
				fmt.Printf("%d \t", input[i][j])
			}

			if j == -1 {
				snake = 1
			}
		}

		i += 1
	}
}

func main() {
	PrintInSnakePattern([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	})
}
