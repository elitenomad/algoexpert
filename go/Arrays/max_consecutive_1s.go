package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Consecutive1s(input []int) int { // O(N)
	maximum := 0
	frequency := 0

	for i := 0; i < len(input); i++ {
		if input[i] == 1 {
			frequency += 1
		} else {
			maximum = max(maximum, frequency)
			frequency = 0
			continue
		}
	}

	return maximum
}

func main() {
	a := []int{1, 0, 1, 1, 1, 1, 0, 1, 1}
	fmt.Println(Consecutive1s(a))
}
