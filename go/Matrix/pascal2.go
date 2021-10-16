package main

import (
	"fmt"
	"strconv"
)

func PascalTriangle(rowthIndex int) []int {
	result := make([]int, rowthIndex+1)
	triangle := make([][]int, rowthIndex+1)

	for i := 0; i < rowthIndex+1; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][len(triangle[i])-1] = 1, 1

		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}

	for i := 0; i < len(result); i++ {
		result[i] = triangle[rowthIndex][i]
	}

	return result
}

func GetNumInPascalTriangle(row, column int, cache *map[string]int) int {
	if row == 0 || column == 0 || row == column {
		return 1
	}

	key := strconv.Itoa(row) + strconv.Itoa(row)

	if _, found := (*cache)[key]; found {
		return (*cache)[key]
	}

	(*cache)[key] = GetNumInPascalTriangle(row-1, column-1, cache) + GetNumInPascalTriangle(row-1, column, cache)

	return (*cache)[key]
}

func FetchValuesOfOthIndex(index int) []int {
	cache := map[string]int{}
	result := make([]int, index+1)
	for i := 0; i < index+1; i++ {
		result[i] = GetNumInPascalTriangle(index, i, &cache)
	}

	return result
}

func main() {
	fmt.Println(FetchValuesOfOthIndex(4))
}
