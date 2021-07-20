package main

import "fmt"

func merge_two_sorted_arrays(a, b []int) []int {
	result := []int{}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j = j + 1
		}
	}

	for i < len(a) {
		result = append(result, a[i])
		i++
	}

	for j < len(b) {
		result = append(result, b[j])
		j++
	}

	return result
}

func sort_sub_arrays_sorted(input []int, low, mid, high int) []int {
	s1 := input[low : mid+1]
	s2 := input[mid+1 : high+1]
	result := merge_two_sorted_arrays(s1, s2)

	return result
}

func MergeSort(a []int) []int {
	return MergeSortHelper(a, 0, len(a)-1)
}

func MergeSortHelper(a []int, low, high int) []int {
	result := []int{}

	if high > low {
		mid := (low + high) / 2
		MergeSortHelper(a, low, mid)
		MergeSortHelper(a, mid+1, high)
		result = sort_sub_arrays_sorted(a, low, mid, high)
	}

	return result
}

func main() {
	a := []int{10, 15, 20, 40, 8, 11, 55}
	fmt.Println(MergeSort(a))
}
