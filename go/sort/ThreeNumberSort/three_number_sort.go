package main

import (
	"fmt"
)

func ThreeNumberSort(array []int, order []int) []int {
	next_index := 0

	for i := 0; i < len(order); i++ {
		var temp []int

		for j := next_index; j < len(array); j++ {
			if order[i] == array[j] {
				temp = append(temp, j)
			}
		}
		array, next_index = swap(array, temp, next_index)
	}

	return array
}

func swap(array []int, temp []int, next_index int) ([]int, int) {
	i := 0
	for i < len(temp) {
		t_1 := array[next_index]
		t_2 := temp[i]

		array[next_index] = array[t_2]
		array[t_2] = t_1
		next_index += 1
		i += 1
	}

	return array, next_index
}

func main() {
	i := []int{1, 0, 0, -1, -1, 0, 1, 1}
	j := []int{0, 1, -1}

	fmt.Println(ThreeNumberSort(i, j))
}
