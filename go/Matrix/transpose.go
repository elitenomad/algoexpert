package main

import "fmt"

func Transpose(input [][]int) [][]int {
	var rows = len(input)
	var cols = len(input[0])
	temp := make([][]int, rows)
	for i := range temp {
		temp[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			temp[i][j] = input[j][i]
		}
	}

	return temp
}

func TransposeII(input [][]int) [][]int {
	var rows = len(input)
	var cols = len(input[0])

	for i := 0; i < rows; i++ {
		for j := i + 1; j < cols; j++ {
			input[i][j], input[j][i] = input[j][i], input[i][j]
		}
	}

	return input
}

func main() {
	fmt.Println(TransposeII([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}))
}
