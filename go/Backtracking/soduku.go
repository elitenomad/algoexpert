package main

import (
	"fmt"
	"math"
)

func Soduku(mesh [][]int, N int) bool {
	i := 0
	j := 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if mesh[i][j] == 0 {
				break
			}
		}
	}

	if i == N && j == N {
		return true
	}

	for n := 1; n < N; n++ {
		if IsSafe(mesh, i, j, N, n) {
			mesh[i][j] = n
			if Soduku(mesh, N) {
				return true
			}
			mesh[i][j] = 0
		}
	}

	return false
}

// Number in a row shouldnt be same
// Number in the column shouldnt be same
// Number in N/2 Grid shuldn't be same
func IsSafe(mesh [][]int, i, j, N, num int) bool {
	for k := 0; k < N; k++ {
		if mesh[k][j] == N || mesh[i][k] == N {
			return false
		}
	}

	sqrt := int(math.Sqrt(float64(N)))
	boxRowStart := i - i%sqrt
	boxColStart := j - j%sqrt

	for r := boxRowStart; r < boxRowStart+sqrt; r++ {
		for d := boxColStart; d < boxColStart+sqrt; d++ {
			if mesh[r][d] == num {
				return false
			}
		}
	}

	return true
}

func main() {
	a := [][]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}
	N := len(a)

	if Soduku(a, N) {
		print(a, N)
	} else {
		fmt.Println("No solution")
	}
}
