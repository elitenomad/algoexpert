package main

import "fmt"

func CycleSortDistinct(input []int) []int {

	for i := 0; i < len(input)-1; i++ {
		// Fix the position and element
		pos := i
		element := input[i]

		// Find the other elements smaller than element
		for j := i + 1; j < len(input); j++ {
			if input[j] < element {
				pos += 1
			}
		}

		// Swap the element onto the position count
		// In the input array
		element, input[pos] = input[pos], element

		for pos != i {
			pos = i
			for k := i + 1; k < len(input); k++ {
				if input[k] < element {
					pos += 1
				}
			}

			element, input[pos] = input[pos], element
		}
	}

	return input
}

func CycleSortWithDuplicates(input []int) []int {

	writes := 0

	for i := 0; i < len(input)-1; i++ {
		// Fix the position and element
		pos := i
		element := input[i]

		// Find the other elements smaller than element
		for j := i + 1; j < len(input); j++ {
			if input[j] < element {
				pos += 1
			}
		}

		if pos == i {
			continue
		}

		for element == input[pos] {
			pos += 1
		}
		// Swap the element onto the position count
		// In the input array
		if pos != i {
			element, input[pos] = input[pos], element
			writes += 1
		}

		for pos != i {
			pos = i
			for k := i + 1; k < len(input); k++ {
				if input[k] < element {
					pos += 1
				}
			}

			for element == input[pos] {
				pos += 1
			}

			if element != input[pos] {
				element, input[pos] = input[pos], element
				writes += 1
			}
		}
	}

	fmt.Println(writes)
	return input
}

func main() {
	fmt.Println("Cycle Sort Implementation...")
	in := []int{4, 2, 7, 1, 3}
	fmt.Println("Output is  :", CycleSortDistinct(in))
	in = []int{4, 7, 7, 1, 3}
	fmt.Println("Output is  :", CycleSortWithDuplicates(in))
	in = []int{4, 3, 2, 1}
	fmt.Println("Output is  :", CycleSortWithDuplicates(in))
}
