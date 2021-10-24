package main

import "fmt"

func FirstOccurence(input []int, target int) int {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (low + high) / 2
		if input[mid] == target {
			if mid == 0 || input[mid] != input[mid-1] {
				return mid
			} else {
				high = mid - 1
			}
		}

		if input[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func LastOccurence(input []int, target int) int {
	low := 0
	high := len(input) - 1

	for low <= high {
		mid := (low + high) / 2
		if mid == len(input)-1 || input[mid] == target {
			if input[mid] != input[mid+1] {
				return mid
			} else {
				low = mid + 1
			}
		}

		if input[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func SearchRanges(input []int, target int) []int {
	first := FirstOccurence(input, target)
	last := LastOccurence(input, target)

	return []int{first, last}
}

func main() {
	a := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(SearchRanges(a, 8))
}
