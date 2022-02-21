package main

import (
	"fmt"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

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

func StockSpanII(input []int) []int {
	// Use stack to store closest greater element, remove it when its un-necessary
	// If there is previous greater element exist => currentIndex - previousGreaterElement Index
	// if not index + 1
	stack := lls.New()
	result := make([]int, len(input))
	result[0] = 1
	stack.Push(0)
	for i := 1; i < len(input); i++ {
		val, _ := stack.Peek()
		for !stack.Empty() && input[val.(int)] <= input[i] {
			stack.Pop()
			val, _ = stack.Peek()
		}

		if stack.Empty() {
			result[i] = i + 1
		} else {
			temp, _ := stack.Peek()
			result[i] = i - temp.(int)
		}
		stack.Push(i)
	}

	return result
}

func main() {
	// a := []int{13, 15, 12, 14, 16, 8, 6, 4, 10, 30}
	// fmt.Println(StockSpan(a))
	// a = []int{10, 20, 30, 40}
	// fmt.Println(StockSpan(a))
	a := []int{40, 30, 20, 10}
	fmt.Println(StockSpan(a))

	a = []int{40, 30, 20, 10}
	fmt.Println(StockSpanII(a))
}
