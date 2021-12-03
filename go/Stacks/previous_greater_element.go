package main

import "fmt"

// Find the previous Greater element closest by position
func PreviousGreaterElement(input []int) []int {
	result := make([]int, len(input))
	result[0] = -1

	for i := 0; i < len(input); i++ {
		var j int
		for j = i - 1; j >= 0; j-- {
			if input[i] < input[j] {
				fmt.Printf("%d \t", input[j])
				break
			}
		}
		if j == -1 {
			fmt.Printf("%d \t", -1)
		}
	}

	return result
}

func PreviousGreaterElementII(input []int) []int {
	result := make([]int, len(input))
	stack := NewStack()
	stack.Push(input[0])

	for i := 0; i < len(input); i++ {
		for !stack.IsEmpty() && stack.Top() <= input[i] {
			stack.Pop()
		}
		var res int
		if stack.IsEmpty() {
			res = -1
		} else {
			res = stack.Pop()
		}

		fmt.Println("<<", res, ">>")
		stack.Push(input[i])
	}

	return result
}

func main() {
	a := []int{20, 30, 10, 5, 15}
	fmt.Println(PreviousGreaterElement(a))
}
