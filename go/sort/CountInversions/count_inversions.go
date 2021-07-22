package main

import "fmt"

func CountInversionsInEfficient(input []int) int { // O(n ^ 2)
	count := 0

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] > input[j] {
				count += 1
			}
		}
	}

	return count
}

/*
	Merge sort implementation
*/
func CountInversions(array []int) int {
	return CountInversionsHelper(array, 0, len(array)-1)
}

func CountInversionsHelper(array []int, left, right int) int {
	count := 0

	if left < right {
		middle := left + (right-left)/2
		count += CountInversionsHelper(array, left, middle)
		count += CountInversionsHelper(array, middle+1, right)
		count += MergeAndCount(array, left, middle, right)
	}

	return count
}

func MergeAndCount(array []int, left, middle, right int) int {
	n1 := middle - left + 1
	n2 := right - middle

	// Make array of size n1 + n2
	leftie := make([]int, n1)
	rightie := make([]int, n2)

	for i := 0; i < n1; i++ {
		leftie[i] = array[i+left]
	}
	for j := 0; j < n2; j++ {
		rightie[j] = array[middle+1+j]
	}

	i := 0
	j := 0
	k := left

	res := 0 // Count which keeps track of i < j and a[i] > a[j]

	for i < n1 && j < n2 {
		if leftie[i] <= rightie[j] {
			array[k] = leftie[i]
			i++
		} else {
			array[k] = rightie[j]
			j++
			res = res + (n1 - i) // Keeps track of the count
		}
		k++
	}

	for i < n1 {
		array[k] = leftie[i]
		i++
		k++
	}

	for j < n2 {
		array[k] = rightie[j]
		j++
		k++
	}

	return res
}

func main() {
	a := []int{20, 10, 5, 2, 1}
	fmt.Println(CountInversions(a))
}
