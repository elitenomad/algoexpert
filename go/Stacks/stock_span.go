package main

import "fmt"

func StockSpan(input []int) []int {
	result := []int{}

	for i := 0; i < len(input); i++ {
		count := 1
		for j := 0; j < i; j++ {
			if input[i] >= input[j] {
				count += 1
			} else {
				break
			}
		}

		result = append(result, count)
	}

	return result
}

func main() {
	a := []int{13, 15, 12, 14, 16, 8, 6, 4, 10, 30}
	fmt.Println(StockSpan(a))
	a = []int{10, 20, 30, 40}
	fmt.Println(StockSpan(a))
	a = []int{40, 30, 20, 10}
	fmt.Println(StockSpan(a))
}
