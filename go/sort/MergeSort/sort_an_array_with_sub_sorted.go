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

func main() {
	a := []int{10, 15, 20, 40, 8, 11, 55}
	low := 0
	mid := 3
	high := 6

	fmt.Println(sort_sub_arrays_sorted(a, low, mid, high))
}
