package main

import "fmt"

func NQueen(chess [][]int, n int) bool {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	if NQueenRecursive(&chess, &result, n, 0) == false {
		return false
	} else {
		fmt.Println(chess)
		return true
	}
}

func NQueenRecursive(chess *[][]int, result *[][]int, n, col int) bool {
	if col == n {
		return true
	}

	for i := 0; i < n; i++ {
		if IsSafe(*chess, i, col) {
			(*chess)[i][col] = 1
			if NQueenRecursive(chess, result, n, col+1) == true {
				return true
			}
			(*chess)[i][col] = 0
		}
	}
	return false
}

func IsSafe(chess [][]int, row, col int) bool {
	for i := 0; i < col; i++ {
		if chess[row][i] == 1 {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chess[i][j] == 1 {
			return false
		}
	}

	for i, j := row, col; j >= 0 && i < len(chess); i, j = i+1, j-1 {
		if chess[i][j] == 1 {
			return false
		}
	}
	return true
}

func main() {
	chess := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	chess_2 := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	fmt.Println(NQueen(chess, 4))
	fmt.Println(NQueen(chess_2, 3))

}
