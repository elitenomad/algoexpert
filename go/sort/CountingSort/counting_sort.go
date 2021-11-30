package main

import "fmt"

func CountingSort(input []int, k int) []int {
	temp := make([]int, k)

	for i := 0; i < len(input); i++ {
		temp[input[i]] += 1
	}

	index := 0
	for i := 0; i < k; i++ {
		for j := 0; j < temp[i]; j++ {
			input[index] = i
			index += 1
		}
	}

	return input
}

func CountingSortGeneralPurpose(input []int, k int) []int {
	count := make([]int, k)
	fmt.Println(count)
	for i := 0; i < len(input); i++ {
		count[input[i]] += 1
	}
	fmt.Println(count)
	for i := 1; i < k; i++ {
		count[i] = count[i-1] + count[i]
	}
	fmt.Println(count)

	output := make([]int, len(input))
	fmt.Println(output)
	for i := len(input) - 1; i >= 0; i-- {
		output[count[input[i]]-1] = input[i]
		count[input[i]] -= 1
	}

	for i := 0; i < len(output); i++ {
		input[i] = output[i]
	}

	return input
}

func main() {
	fmt.Println("CountingSort Implementation...")
	in := []int{1, 4, 4, 1, 0, 1}
	fmt.Println("Output is  :", CountingSort(in, 5))
	fmt.Println("Output is  :", CountingSortGeneralPurpose(in, 5))

	// in = []int{2, 1, 8, 9, 4}
	// fmt.Println("Output is  :", CountingSort(in, 10))
	// in = []int{4, 3, 2, 1}
	// fmt.Println("Output is  :", CountingSort(in, 5))
}
