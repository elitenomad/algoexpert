package main

import "fmt"

func SecondLargestElementOfArray(input []int) int {
	largest := input[0]
	prevLargest := input[1]

	for i := 1; i < len(input); i++ {
		if largest < input[i] {
			prevLargest = largest
			largest = input[i]
		} else {
			if prevLargest < input[i] {
				prevLargest = input[i]
			}
		}
	}

	return prevLargest
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
	a = []int{3, 1, 6, 5, 10, 9}
	fmt.Println(SecondLargestElementOfArray(a))
}
