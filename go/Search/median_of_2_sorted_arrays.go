package main

import (
	"fmt"
	"math"
)

// Two sorted arrays
func MedianOfTwoSortedArrays(i1, i2 []int) float64 { // O(MAX(I1,I2)), O(I1+I2)
	i := 0
	j := 0
	i3 := []int{}

	for i != len(i1)-1 || j != len(i2)-1 {
		if i1[i] <= i2[j] {
			i3 = append(i3, i1[i])
			i++
		} else {
			i3 = append(i3, i2[j])
			j++
		}
	}

	for i < len(i1) {
		i3 = append(i3, i1[i])
		i++
	}

	for j < len(i2) {
		i3 = append(i3, i2[j])
		j++
	}

	sizeOfI3 := len(i3)
	if sizeOfI3%2 != 0 {
		// For odd numbers we have n/2
		return float64(i3[sizeOfI3/2])
	} else {
		// For even we need to find middle two numbers, n/2 and n/2 -1
		n1 := i3[sizeOfI3/2]
		n2 := i3[sizeOfI3/2-1]
		return float64((n1 + n2)) / 2.0
	}
}

func MedianOfTwoSortedArraysBinarySearch(i1 []int, i2 []int) float64 {
	// Ensure i1 < i2 always for the approach in this method
	// In case input i2 > i1 , we follow the same approach reversing the references
	// where i2 ->  i1 and i1 -> i2
	// We maintain two sets
	// 1 set -> some elements from i1 and some elements from i2
	// 2nd set -> Some elements from i1 and some elements from i2

	return MedianOfTwoSortedArraysBinarySearchHelper(i1, i2, 0, len(i1)-1, 0, len(i2)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func MedianOfTwoSortedArraysBinarySearchHelper(i1 []int, i2 []int, i1_start, i1_end, i2_start, i2_end int) float64 {
	if i1_end-i1_start == 1 && i2_end-i2_start == 1 {
		median := ((float64(max(i1[i1_start], i2[i2_start])) + float64(min(i1[i1_end], i2[i2_end]))) / 2.0)

		return median
	}

	m1_index := (i1_start + i1_end) / 2
	m2_index := (i2_start + i2_end) / 2

	m1 := i1[m1_index]
	m2 := i2[m2_index]

	if m1 == m2 {
		return float64(m1)
	}

	if m1 < m2 { // Narrow down search by eliminating l
		i1_start = m1_index
		i2_end = m2_index
	} else {
		i1_end = m1_index
		i2_start = m2_index
	}

	return MedianOfTwoSortedArraysBinarySearchHelper(i1, i2, i1_start, i1_end, i2_start, i2_end)
}

func MedianOfTwoSortedArraysII(i1, i2 []int) float64 {
	if len(i1) > len(i2) {
		return MedianOfTwoSortedArraysII(i2, i1)
	}

	low := 0
	high := len(i1)

	for low <= high {
		cut1 := (low + high) / 2
		cut2 := (len(i1)+len(i2))/2 - cut1

		var l1 int
		var l2 int
		var r1 int
		var r2 int
		if cut1 == 0 {
			l1 = math.MinInt32
		} else {
			l1 = i1[cut1-1]
		}
		if cut2 == 0 {
			l2 = math.MinInt32
		} else {
			l2 = i2[cut2-1]
		}

		if cut1 == len(i1) {
			r1 = math.MaxInt32
		} else {
			r1 = i1[cut1]
		}

		if cut2 == len(i2) {
			r2 = math.MaxInt32
		} else {
			r2 = i2[cut2]
		}

		if l1 <= r2 && l2 <= r1 {
			if (len(i1)+len(i2))%2 == 0 {
				return (float64(max(l1, l2)) + float64(min(r1, r2))) * 0.5
			} else {
				return (float64(min(r1, r2)))
			}
		} else if l1 > r2 {
			high = cut1 - 1
		} else if l2 > r1 {
			low = cut1 + 1
		}
	}

	return 0.0
}

func main() {
	i1 := []int{10, 20, 30, 40, 50}
	i2 := []int{5, 15, 25, 35, 45}
	// fmt.Println(MedianOfTwoSortedArrays(i1, i2))
	fmt.Println(MedianOfTwoSortedArraysBinarySearch(i1, i2))
}
