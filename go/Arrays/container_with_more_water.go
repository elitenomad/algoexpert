package main

import "fmt"

func maxArea(height []int) int {
	maxArea := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			currentArea := min(height[i], height[j]) * (j - i)
			maxArea = max(maxArea, currentArea)
		}
	}

	return maxArea
}

func maxAreaEfficient(height []int) int {
	maxArea := 0

	i := 0
	j := len(height) - 1
	for i < j {
		currentArea := min(height[i], height[j]) * (j - i)
		maxArea = max(maxArea, currentArea)

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(maxAreaEfficient(a))

	a = []int{5, 4, 3}
	fmt.Println(maxAreaEfficient(a))

	a = []int{2, 0, 2}
	fmt.Println(maxAreaEfficient(a))
}
