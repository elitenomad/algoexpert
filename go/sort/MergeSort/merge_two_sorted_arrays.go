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

func main() {
	a := []int{10, 15, 20}
	b := []int{5, 6, 6, 15}
	fmt.Println(merge_two_sorted_arrays(a, b))
}
