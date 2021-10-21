package main

import "fmt"

func SortedRotatedSearch(input []int, n int) int {
	// Find the mid
	low := 0
	high := len(input) - 1
	mid := (low + high) / 2

	// verify which half of the array is sorted
	if input[mid] == n {
		return mid
	}

	if input[mid] > input[0] {
		// Left half is sorted,
		// Binary search in left half
		h_high := mid - 1
		for low <= h_high {
			mid = (low + h_high) / 2
			if input[mid] == n {
				return mid
			}

			if input[mid] < n {
				low = mid + 1
			} else {
				h_high = mid - 1
			}
		}
		// Linear search in right half
		low = mid + 1
		high = len(input) - 1
		for i := low; i < high; i++ {
			if input[i] == n {
				return i
			}
		}

	}

	if input[mid] < input[0] {
		// Linear search in right half.
		low = 0
		high = mid - 1
		for i := low; i < high; i++ {
			if input[i] == n {
				return i
			}
		}
		// Right half is sorted
		// Binary search in Right half
		low = mid + 1
		high = len(input) - 1
		for low <= high {
			mid = (low + high) / 2
			if input[mid] == n {
				return mid
			}

			if input[mid] < n {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func SortedRotatedSearchImprovise(input []int, n int) int {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (low + high) / 2
		if input[mid] == n {
			return mid
		}

		if input[low] <= input[mid] {
			//Left half sorted
			if n >= input[low] && n < input[mid] {
				// No need to check for <= input[mid]as we already did at line: 79
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			//Right half sorted
			if n > input[mid] && n <= input[len(input)-1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func main() {
	a := []int{100, 500, 10, 20, 30, 40, 50}
	fmt.Println(SortedRotatedSearchImprovise(a, 50))
}
