package main

import (
	"fmt"
	"sort"
)

func MergeNaive(nums1 []int, m int, nums2 []int, n int) []int {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1) // O((M+N)Log(M+N))
	return nums1
}

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	numsCopy := make([]int, m+n)
	copy(numsCopy, nums1)

	i := 0 // nums1
	j := 0 // nums2
	k := 0 // numsCopy

	for k < m+n {
		fmt.Println(i, j, k)
		if j >= n || (i < m && numsCopy[i] <= nums2[j]) {
			nums1[k] = numsCopy[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}

		k++
	}

	return nums1
}

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	m := 3
	b := []int{2, 5, 6}
	n := 3
	// fmt.Println(MergeNaive(a, m, b, n))
	fmt.Println(Merge(a, m, b, n))
}
