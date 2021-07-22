package main

import "fmt"

// func merge_two_sorted_arrays(a, b []int) []int {
// 	result := []int{}

// 	i := 0
// 	j := 0

// 	for i < len(a) && j < len(b) {
// 		if a[i] <= b[j] {
// 			result = append(result, a[i])
// 			i++
// 		} else {
// 			result = append(result, b[j])
// 			j = j + 1
// 		}
// 	}

// 	for i < len(a) {
// 		result = append(result, a[i])
// 		i++
// 	}

// 	for j < len(b) {
// 		result = append(result, b[j])
// 		j++
// 	}

// 	return result
// }

// func sort_sub_arrays_sorted(input []int, low, mid, high int) []int {
// 	s1 := input[low : mid+1]
// 	s2 := input[mid+1 : high+1]
// 	result := merge_two_sorted_arrays(s1, s2)
// 	fmt.Println(result)
// 	return result
// }

// func MergeSort(a []int) []int {
// 	return MergeSortHelper(a, 0, len(a)-1)
// }

// func MergeSortHelper(a []int, low, high int) []int {
// 	if len(a) <= 1 {
// 		return a
// 	}
// 	result := []int{}

// 	if high > low {
// 		mid := (low + high) / 2
// 		MergeSortHelper(a, low, mid)
// 		MergeSortHelper(a, mid+1, high)
// 		doMerge(a, low, mid, high)
// 	}

// 	return result
// }

func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	middle := len(array) / 2
	left := MergeSort(array[:middle])
	right := MergeSort(array[middle:])

	return merge_two_sorted_arrays(left, right)
}

func merge_two_sorted_arrays(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func main() {
	a := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(MergeSort(a))
}
