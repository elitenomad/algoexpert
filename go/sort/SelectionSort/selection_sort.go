package main

import "fmt"

func SelectionSort(array []int) []int {
	currentIndex := 0

	//Two Loops
	// Outer one -> Loops through Entire array
	// Inner One -> Loops through Subset of array
	for currentIndex < len(array)-1 {
		smallestIndex := currentIndex
		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] > array[i] {
				smallestIndex = i
			}
		}
		array[currentIndex], array[smallestIndex] = array[smallestIndex], array[currentIndex]
		currentIndex = currentIndex + 1
	}

	return array
}

func main() {
	i := []int{3, 1, 2, 4, 5, 6}
	fmt.Println(SelectionSort(i))
}
