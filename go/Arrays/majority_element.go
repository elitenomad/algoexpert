package main

import "fmt"

func MajorityElement(input []int) []int {
	size := len(input)
	majority := (size / 2) + 1
	result := []int{}
	h := map[int][]int{}

	for i, value := range input {
		h[value] = append(h[value], i)
	}

	for key, indexes := range h {
		if len(indexes) >= majority {
			result = h[key]
		}
	}

	return result
}

/*
	- Find the Majority
	- Print the indexes of the Majority
*/

func FindMajority(input []int) int { //{8, 3, 8, 1, 8, 2, 8, 9, 8}
	majority := 0
	count := 1

	for i := 1; i < len(input); i++ {
		if input[majority] == input[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			majority = i
			count = 1
		}
	}

	return input[majority]
}

func MajorityElementEfficient(input []int) int {
	majority := FindMajority(input)
	count := 0

	for _, val := range input {
		if majority == val {
			count++
		}
	}

	if count <= len(input)/2 {
		return -1
	}

	return count
}

func main() {
	a := []int{8, 3, 4, 8, 8}
	fmt.Println(MajorityElement(a))
	fmt.Println(MajorityElementEfficient(a))
	a = []int{8, 3, 8, 1, 8, 2, 8, 9}
	fmt.Println(MajorityElement(a))
	fmt.Println(MajorityElementEfficient(a))
}
