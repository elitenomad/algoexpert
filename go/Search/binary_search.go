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

func BinarySearchRecursive(input []int, low, high, element int) int {
	mid := (low + high) / 2

	if input[mid] == element {
		return mid
	}

	if low <= high {
		if input[mid] < element {
			return BinarySearchRecursive(input, mid+1, high, element)
		} else if input[mid] > element {
			return BinarySearchRecursive(input, low, mid-1, element)
		}
		mid = (low + high) / 2
	}

	return -1
}

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(BinarySearchIterative(a, 50))
	fmt.Println(BinarySearchIterative(a, 560))

	low := 0
	high := len(a) - 1
	fmt.Println(BinarySearchRecursive(a, low, high, 50))
	fmt.Println(BinarySearchRecursive(a, low, high, 560))
}
