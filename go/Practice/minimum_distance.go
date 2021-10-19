// Minimum distance between two numbers

package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func MinimumDistanceBetweenTwoPoints(input []int, x, y int) int {
	result := 100000000
	posX := -1
	posY := -1

	for i := 0; i < len(input); i++ {
		if input[i] == x {
			posX = i
		} else if input[i] == y {
			posY = i
		}

		fmt.Println(posX, posY)
		if posX >= 0 && posY >= 0 {
			result = min(result, abs(posX-posY))
		}
	}

	return result // Large number means its not present
}

func main() {
	a := []int{1, 2, 3, 2}
	fmt.Println(MinimumDistanceBetweenTwoPoints(a, 1, 2))
	a = []int{86, 39, 90, 67, 84, 66, 62}
	fmt.Println(MinimumDistanceBetweenTwoPoints(a, 42, 12))
}
