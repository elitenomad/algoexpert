package main

import "fmt"

func PeakElement(input []int) { //O(N)

	for i := 1; i < len(input)-1; i++ {
		if i == 1 && input[i-1] >= input[i] {
			fmt.Println(input[i-1])
		}

		if i == len(input)-1 && input[len(input)-1] >= input[len(input)-2] {
			fmt.Println(input[len(input)-1])
		}

		if input[i] >= input[i-1] && input[i] >= input[i+1] {
			fmt.Println(input[i])
		}
	}
}

// If we take this Approach
// And if there are multiple peak elements
// We might end up printing not all the peak Elements
// Available
func PeakElementLogOfN(input []int) int {
	low := 0
	high := len(input) - 1
	n := len(input) - 1

	for low <= high {
		mid := (low + high) / 2
		// Two conditions
		// mid == 0
		// mid == n - 1
		if ((mid == 0) || input[mid-1] <= input[mid]) && ((mid == n) || (input[mid+1] <= input[mid])) {
			return input[mid]
		}

		if mid > 0 && input[mid-1] > input[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	a := []int{5, 10, 20, 15, 7}
	PeakElement(a)
	fmt.Println(PeakElementLogOfN(a))
	fmt.Println()
	a = []int{80, 70, 60}
	PeakElement(a)
	fmt.Println(PeakElementLogOfN(a))
	fmt.Println()

	a = []int{10, 20, 15, 5, 23, 80, 70}
	PeakElement(a)
	fmt.Println(PeakElementLogOfN(a))
	fmt.Println()
	a = []int{10, 20}
	fmt.Println(PeakElementLogOfN(a))
}
