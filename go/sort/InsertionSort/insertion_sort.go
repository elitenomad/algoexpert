package main

import "fmt"

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1

		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}

		array[j+1] = key
	}
	return array
}

func main() {
	input := []int{3, 4, 2, 1, 5}
	fmt.Println(InsertionSort(input))
}
