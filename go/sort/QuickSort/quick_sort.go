package main

import "fmt"

// Worst Case : O(n^2)
// Best and Avg case: O(NLog(n))
// Space complexity : O(Log(n))
func QuickSort(array []int) []int {
	return QuickSortHelper(array, 0, len(array)-1)
}

func QuickSortHelper(array []int, start, end int) []int {
	if start >= end {
		return array
	}

	pivot := start
	left := start + 1
	right := end

	for right >= left {
		if array[left] > array[pivot] && array[right] < array[pivot] {
			array[left], array[right] = array[right], array[left]
		}

		if array[left] <= array[pivot] {
			left += 1
		}

		if array[right] >= array[pivot] {
			right -= 1
		}
	}

	array[pivot], array[right] = array[right], array[pivot]

	leftArrayIsSmaller := (right - 1 - start) < (end - right - 1)

	if leftArrayIsSmaller {
		QuickSortHelper(array, start, right-1)
		QuickSortHelper(array, right+1, end)
	} else {
		QuickSortHelper(array, right+1, end)
		QuickSortHelper(array, start, right-1)
	}

	return array
}

func main() {
	i := []int{3, 1, 2, 4, 5, 6}
	fmt.Println(QuickSort(i))
}
