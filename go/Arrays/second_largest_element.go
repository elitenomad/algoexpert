package main

import "fmt"

func SecondLargestElementOfArray(input []int) int {
	result := -1
	largest := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[largest] {
			result = largest
			largest = i
		}

		if input[i] == input[largest] {
			continue
		}

		if input[i] < input[largest] {
			if result == -1 {
				result = i
			}

			if input[i] < input[result] {
				// Ignore
			} else {
				result = i
			}
		}
	}

	return result
}

/*
	Another Naive approach is to
		- First find the largest number
		- Loop through all the number again except for the largest number
		- Find the largest number in the remaining elements
*/

func main() {
	a := []int{1, 3, 9, 5, 10, 18, 17, 14} // 18,17,14,10
	fmt.Println(SecondLargestElementOfArray(a))
}
