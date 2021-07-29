package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func isArraysOverlap(i1, i2 []int) bool {
	// Find the maximum of first element of two arrays
	max_of_first_elements := max(i1[0], i2[0])

	// Initialize the array of length 2 to compare it to max element found above
	array_to_be_verifed := make([]int, 2)

	// Prepare array to be verified
	if max_of_first_elements == i1[0] {
		copy(array_to_be_verifed, i2)
	}

	if max_of_first_elements == i2[0] {
		copy(array_to_be_verifed, i1)
	}

	// validate if max value collected is in b/w array to be verifed
	if array_to_be_verifed[0] <= max_of_first_elements && array_to_be_verifed[1] > max_of_first_elements {
		return true
	}

	// If the above conditon doesn't pass implies that arrays are not overlapping.
	return false
}

type Interval struct {
	start int
	end   int
}

type Intervals []Interval

func (s Intervals) Len() int      { return len(s) }
func (s Intervals) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByStart struct{ Intervals }

func (s ByStart) Less(i, j int) bool { return s.Intervals[i].start < s.Intervals[j].start }

func MergeOverlappingIntervals(intervals Intervals) {
	sort.Sort(ByStart{intervals})
	res := 0

	for i := 0; i < len(intervals); i++ {
		if intervals[res].end >= intervals[i].start {
			intervals[res].end = max(intervals[res].end, intervals[i].end)
			intervals[res].start = max(intervals[res].start, intervals[i].start)
		} else {
			res++
			intervals[res] = intervals[i]
		}
	}

	for i := 0; i <= res; i++ {
		fmt.Println("[", intervals[i].start, intervals[i].end, "]")
	}
}

func main() {
	// a := []int{5, 10}
	// b := []int{1, 7}
	// fmt.Println(isArraysOverlap(a, b))
	// a = []int{10, 20}
	// b = []int{5, 15}
	// fmt.Println(isArraysOverlap(a, b))
	// a = []int{10, 20}
	// b = []int{500, 505}
	// fmt.Println(isArraysOverlap(a, b))

	a := Intervals{
		Interval{5, 10},
		Interval{3, 15},
		Interval{18, 30},
		Interval{2, 7},
	}

	MergeOverlappingIntervals(a)
}
