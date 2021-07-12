package main

import "fmt"

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		temp := array[i]
		j := i - 1

		for j >= 0 {
			if array[j] > temp {
				array[j+1] = array[j]
				j = j - 1
			} else {
				break
			}
		}
		array[j+1] = temp
	}
	return array
}

func main() {
	input := []int{3, 4, 2, 1, 5}
	fmt.Println(InsertionSort(input))
}
