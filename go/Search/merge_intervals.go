package main

import "fmt"

// First question to ask is if the arrays are sorted by first element
// If not sort them
func MergeOverlappingIntervals(intervals [][]int) [][]int {

	// result := [][]int{}
	j := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[j][1] >= intervals[i][0] {
			first := min(intervals[j][0], intervals[i][0])
			second := max(intervals[j][1], intervals[i][1])

			intervals[j][0] = first
			intervals[j][1] = second
			// result = append(result, []int{first, second})
		} else {
			j += 1
			intervals[j] = intervals[i]
			// result = append(result, intervals[i])
		}
	}

	return intervals[0 : j+1]
}

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

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(MergeOverlappingIntervals(intervals))
}
