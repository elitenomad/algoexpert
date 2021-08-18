package main

import "fmt"

// Return the index of n in input array - Last occurence
func SortedLastOccurence(input []int, n int) int {
	result := -1

	for i, val := range input {
		if val == n {
			result = i
		}
	}

	return result
}

func SortedLastOccurenceEff(input []int, n int) int { // O(Logn)
	low := 0
	high := len(input) - 1
	mid := (low + high) / 2

	for low <= high {
		if input[mid] == n {
			if mid == high || input[mid] != input[mid+1] {
				return mid
			} else {
				low = mid + 1
			}
			// Below is one approach
			// temp := mid
			// realMid := mid
			// for temp <= high {
			// 	if input[temp] == input[high] {
			// 		realMid = temp
			// 	}
			// 	temp++
			// }

			// return realMid
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
	fmt.Println(SortedLastOccurenceEff(a, 15))
	a = []int{5, 5, 5}
	fmt.Println(SortedLastOccurenceEff(a, 5))
}
