package main

import "fmt"

func PascalTriangle(numOfRows int) [][]int {
	result := make([][]int, numOfRows)

	for i := 0; i < len(result); i++ {
		result[i] = make([]int, i+1)
		result[i][0], result[i][len(result[i])-1] = 1, 1

		for j := 1; j < len(result[i])-1; j++ {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
		fmt.Println(result[i])
	}

	return result
}

func main() {
	PascalTriangle(5)
}
