package main

import "fmt"

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
	return -1
}

func main() {
	i1 := []int{10, 20, 30, 40, 50}
	i2 := []int{5, 15, 25, 35, 45}
	fmt.Println(MedianOfTwoSortedArrays(i1, i2))
}
