package main

import "fmt"

func CountOccurencesinSortedArray(input []int, n int) int {
	result := 0

	low := 0
	high := len(input) - 1
	mid := (low + high) / 2

	for low <= high {
		if input[mid] == n {
			result += 1
			temp_l := low
			temp_h := high

			for temp_l < mid {
				if input[temp_l] == n {
					result += 1
				}

				temp_l++
			}

			for mid < temp_h {
				if input[temp_h] == n {
					result += 1
				}

				temp_h--
			}

			return result
		}

		if input[mid] < n {
			low = mid + 1
		} else if input[mid] > n {
			high = mid - 1
		}

		mid = (low + high) / 2
	}

	return result
}

// Another approach is
// Use two Binary searches
// First Occurence of an element
// last occurent of an element
// Last-First + 1

func main() {
	a := []int{10, 20, 20, 20, 20, 30, 30}
	fmt.Println(CountOccurencesinSortedArray(a, 30))
}
