package main

import "fmt"

// Return the index of n in input array
func SortedFirstOccurence(input []int, n int) int {
	result := -1

	for i, val := range input {
		if val == n {
			result = i
			break
		}
	}

	return result
}

func SortedFirstOccurenceEff(input []int, n int) int { // O(Logn)
	low := 0
	high := len(input) - 1
	mid := (low + high) / 2

	for low <= high {
		if input[mid] == n {
			if mid == 0 || (input[mid-1] != input[mid]) {
				return mid
			} else {
				high = mid - 1
				// for temp <= mid-1 {
				// 	if input[temp] == n {
				// 		return temp
				// 	}
				// 	temp++
				// }
			}
		}

		if input[mid] < n {
			low = mid + 1
		} else if input[mid] > n {
			high = mid - 1
		}

		mid = (low + high) / 2
	}

	return -1
}

func main() {
	a := []int{5, 10, 10, 15, 15}
	fmt.Println(SortedFirstOccurenceEff(a, 15))
	a = []int{5, 5, 5}
	fmt.Println(SortedFirstOccurenceEff(a, 5))
}
