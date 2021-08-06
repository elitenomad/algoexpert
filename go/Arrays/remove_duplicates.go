// Remove duplicate from sorted array

package main

import "fmt"

func RemoveDuplicates(input []int) int { // O(n) for both T and S
	temp := []int{}

	// {2,2,3,4,4,4,4,5}
	previous := input[0]
	temp = append(temp, previous)

	for i := 1; i < len(input); i++ {
		if input[i] != previous {
			temp = append(temp, input[i])
			previous = input[i]
		}
	}

	return len(temp)
}

func RemoveDuplicatesEfficient(input []int) int { // O(n) T and O(1) S
	result := 1

	for i := 1; i < len(input)-1; i++ {
		if input[i] != input[result-1] {
			// If you want to change the array
			input[result] = input[i]
			result++
		}
	}

	input = input[:result]
	fmt.Println(input)

	// Goal is to return the length of an array without
	// Duplicates
	return result
}

func main() {
	a := []int{2, 2, 3, 4, 4, 4, 4, 5}
	fmt.Println(RemoveDuplicates(a))
	fmt.Println(RemoveDuplicatesEfficient(a))
}
