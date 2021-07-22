package main

import "fmt"

func NaiveParition(input []int, pivot int) int {
	return NaiveParitionHelper(input, 0, len(input)-1, pivot)
}

func NaiveParitionHelper(input []int, low, high, pivot int) int { //O(n) T and O(n) Ss
	//Swap the pivot with last element
	input[high], input[pivot] = input[pivot], input[high]
	pivot = high
	fmt.Println(input)

	// Create an auxillary array with len of high - low + 1 or high + low
	temp := make([]int, high-low+1)
	index := 0

	// Copy small elements onto temp array
	for i := low; i <= high; i++ {
		if input[i] <= input[high] {
			temp[index] = input[i]
			index++
		}
	}
	fmt.Println(temp)
	// Find the Index of the pivot element when moves in equality condition
	pivot = low + index - 1

	// Copy large elements onto temp array
	for i := low; i <= high; i++ {
		if input[i] > temp[pivot] {
			temp[index] = input[i]
			index++
		}
	}

	copy(input, temp)
	fmt.Println(input)

	return pivot
}

func main() {
	a := []int{3, 8, 6, 12, 10, 7}
	fmt.Println(NaiveParition(a, 2))
}
