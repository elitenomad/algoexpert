package main

import (
	"fmt"
)

func contains(ints []int, val int) bool {
	for _, v := range ints {
		if v == val {
			return true
		}
	}

	return false
}

/*
	Intersection of Two SORTED || UNSORTED arrays
	[1,1,3,3,3] & [1,1,1,1,1,3,5,7] = [1,3]
*/
func Intersection(a []int, b []int) []int {
	var result []int

	for i := 0; i < len(a); i++ { // O(len(A)*len(B))
		if i > 0 && a[i] == a[i-1] {
			continue
		}

		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				result = append(result, b[j])
				break
			}
		}
	}

	return result
}

/*
	Intersection of two SORTED arrays
	Using merge sort , merge functionality
*/
func IntersectionOfTwoSortedArrays(a, b []int) []int {
	result := []int{}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}

		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			result = append(result, a[i])
			i++
			j++
		}
	}

	return result
}

func main() {
	a := []int{1, 1, 3}
	b := []int{1, 1, 1, 1, 1, 3, 7}
	fmt.Println(Intersection(a, b))
	fmt.Println(IntersectionOfTwoSortedArrays(a, b))
	a = []int{1, 1, 2, 2}
	b = []int{2, 2}
	fmt.Println(Intersection(a, b))
}
