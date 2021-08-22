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

// func MajorityElementUsingSorting(input []int) int {
// 	sort.Ints(input)
// 	maxElement := input[0]
// 	maxIndex := 0

// 	for i := 1; i < len(input); i++ {
// 		maxCount := 1

// 		/*
// 			While i + 1 < N && input[i] == input[i+1]{
// 				count = count + 1
// 			}

// 			if  count > N/2 {
// 				print maxElement
// 				brea
// 			}
// 		*/
// 		if input[i] == maxElement {
// 			maxElement = input[i]
// 			maxCount += 1
// 		}

// 	}
// }

/*
	- Find the Majority
	- Print the Majority
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

func FindMaxOccuringElem(input []int) int { //{7,8,7,7,8,7,7,9}
	index := 0
	count := 1

	for i := 1; i < len(input); i++ {
		if input[index] == input[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			index = i
			count = 1
		}
	}

	return input[index]
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
	a = []int{7, 8, 7, 7, 8, 7, 7, 9}
	fmt.Println(FindMaxOccuringElem(a))
}
