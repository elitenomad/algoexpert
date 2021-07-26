package main

import (
	"fmt"
	"sort"
)

/*
	One approach would be Sort the
	given array using Selection Sort and when
	K matches the loop count return the result

	Best case: O(1)
	Average Case: O(n^2)
	Worst Case: O(n^2)
*/
func KthSmallestElement(input []int, k int) int {
	j := 0

	for j < len(input)-1 {
		smallestIndex := j

		for i := j + 1; i < len(input); i++ {
			if input[smallestIndex] > input[i] {
				smallestIndex = i
			}
		}
		input[smallestIndex], input[j] = input[j], input[smallestIndex]
		if j == k {
			break
		}
		j++
	}

	return input[j-1]
}

/*
	Another approach would be use the sort function
	of the language and the find k-1 element in the array
	Sort generally Time complexity would be - O(Nlog(N))
*/
func KthSmallestElementUsingGoSort(input []int, k int) int {
	soretedArray := sort.IntSlice(input)
	return soretedArray[k-1]
}

/*
	Another approach would be to use Lumoto Partition
	approach to figure out kth Smallest Number
*/
func Partition(input []int, start, end int) int {
	i := start - 1

	// Find the length of an array
	// len := len(input) - 1
	// To start with we will consider our pivot as the last element in the array
	pivot := end

	// Loop through the array
	for j := 0; j < end; j++ {
		if input[j] < input[pivot] { // if an element < pivot increment i and swap values
			i++
			input[j], input[i] = input[i], input[j]
		}
	}
	input[i+1], input[end] = input[end], input[i+1]
	return i + 1
}

func KthSmallestElementUsingLumotoPartition(input []int, n, k int) int {
	l := 0
	r := len(input) - 1

	for l < r {
		p := Partition(input, l, r)
		if p == k-1 {
			return input[p]
		} else if p > k-1 {
			r = p - 1
		} else {
			r = p + 1
		}
	}

	return -1
}

func main() {
	a := []int{10, 5, 30, 2}
	k := 2
	fmt.Println(KthSmallestElement(a, k))
	fmt.Println(KthSmallestElementUsingGoSort(a, k))
	fmt.Println(KthSmallestElementUsingLumotoPartition(a, len(a), k))
	a = []int{30, 20, 5, 10, 8}
	k = 4
	fmt.Println(KthSmallestElement(a, k))
	fmt.Println(KthSmallestElementUsingGoSort(a, k))
	fmt.Println(KthSmallestElementUsingLumotoPartition(a, len(a), k))
}
