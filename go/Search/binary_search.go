package main

import "fmt"

// Return index of the element if present
// If not return -1
func BinarySearchIterative(input []int, element int) int {
	low := 0
	high := len(input) - 1
	mid := (low + high) / 2

	for low <= high {
		// Goal is to change low and high
		// Based on the value
		if input[mid] == element {
			return mid
		}

		if input[mid] < element {
			low = mid + 1
		} else if input[mid] > element {
			high = mid - 1
		}

		mid = (low + high) / 2
	}

	return -1
}

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(BinarySearchIterative(a, 50))
	fmt.Println(BinarySearchIterative(a, 560))
}
