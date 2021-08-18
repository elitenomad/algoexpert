package main

import "fmt"

func CountOnesInSortedArray(arr []int) int {
	low := 0
	high := len(arr) - 1

	mid := (low + high) / 2
	count := 0

	for low <= high {
		if arr[mid] == 1 {
			count += (mid - low) + 1
			low = mid + 1
		} else {
			high = mid - 1
		}

		mid = (low + high) / 2
	}

	return count
}

func FindFirstOccurenceOf1(input []int) int {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (low + high) / 2

		if input[mid] == 0 {
			low = mid + 1
		} else {
			if mid == 0 || input[mid] != input[mid-1] {
				return mid
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

func Count1sInSortedArray(input []int, n int) int {
	result := FindFirstOccurenceOf1(input)
	return len(input) - result
}

func main() {
	a := []int{0, 0, 1, 1, 1, 1, 1}
	fmt.Println(Count1sInSortedArray(a, 1))
}
