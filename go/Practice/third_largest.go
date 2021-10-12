package main

import (
	"fmt"
	"sort"
)

func ThirdLargest(input []int) int { //O(NLOGN)
	sort.Ints(input)

	return input[2]
}

// given all the values in an array are distinct
func ThirdLargestII(input []int) int {
	largest := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[largest] {
			largest = i
		}
	}

	//swap largest with first element
	input[0], input[largest] = input[largest], input[0]
	second_largest := input[1]
	third_largest := -1
	fmt.Println(input)
	for i := 2; i < len(input); i++ {
		if input[i] > second_largest {
			third_largest = second_largest
			second_largest = input[i]
		} else {
			if input[i] > third_largest {
				third_largest = input[i]
			}
		}
	}

	return third_largest
}

func ThirdLargestIII(input []int) int {
	if len(input) < 3 {
		return -1
	}

	first, second, third := input[0], -1, -1

	for i := 1; i < len(input); i++ {
		if input[i] > first {
			third = second
			second = first
			first = input[i]
		} else if input[i] > second {
			third = second
			second = input[i]
		} else if input[i] > third {
			third = input[i]
		}
	}

	return third
}

func main() {
	a := []int{2, 4, 1, 10, 5} // {5,4,1,3,2}
	// fmt.Println(ThirdLargestII(a))
	fmt.Println(ThirdLargestIII(a))

}
