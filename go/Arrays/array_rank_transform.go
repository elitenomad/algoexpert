package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	temp := make([]int, len(arr)) // Space complexity - O(Size of Array)
	copy(temp, arr)
	sort.Ints(temp) // Sorts the temporary array - TC O(NLogN)

	h := map[int]int{}
	rank := 1
	for i := 0; i < len(temp); i++ { // used a hashing strategy to assign Ranks - TC O(N)
		if _, found := h[temp[i]]; !found {
			h[temp[i]] = rank
			rank++
		}
	}

	for i := 0; i < len(temp); i++ { // TC O(N)
		temp[i] = h[arr[i]]
	}

	return temp
}

func main() {
	a := []int{40, 10, 30, 20}
	fmt.Println(arrayRankTransform(a))
}
