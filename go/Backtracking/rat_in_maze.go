package main

import "fmt"

func RatInMaze(maze [][]int) bool {
	N := len(maze)

	// Initialize 2D matrix
	result := make([][]int, N)
	for i := 0; i < len(maze); i++ {
		result[i] = make([]int, N)
	}
	fmt.Println(result)
	if RatInMazeRecursive(maze, &result, 0, 0, N) == false {
		return false
	} else {
		fmt.Println(result)
		return true
	}
}

func RatInMazeRecursive(maze [][]int, result *[][]int, i, j, N int) bool {
	if i == N-1 && j == N-1 { // When you reach the cheese
		return true
	}

	if IsSafe(maze, i, j, N) == true {
		(*result)[i][j] = 1
		if RatInMazeRecursive(maze, result, i, j+1, N) == true {
			return true
		} else if RatInMazeRecursive(maze, result, i+1, j, N) == true {
			return true
		}
		(*result)[i][j] = 0
	}

	return false
}

// If i and j are in the Limits (i >= 0 and i < N , j >= 0 and j < N)
// in the maze i and j reference to 1 (then rat can step in)
func IsSafe(maze [][]int, i, j, N int) bool {
	return i < N && j < N && maze[i][j] == 1
}

func main() {
	maze := [][]int{
		{1, 0, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	fmt.Println(RatInMaze(maze))
}
