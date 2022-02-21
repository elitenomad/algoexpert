package main

import "fmt"

func NextGreaterElementNaive(array []int) {
	for i := 0; i < len(array); i++ {
		var j int
		for j = i + 1; j < len(array); j++ {
			if array[j] > array[i] {
				fmt.Println("<<", array[j], ">>")
				break
			}
		}

		if j == len(array) {
			fmt.Println("<<", -1, ">>")
		}
	}
}

// Thing to note is that it will display the elements in
// the reverse order
func NextGreaterElementEfficient(array []int) {
	stack := NewStack() // 5, 15, 10, 8, 6, 12, 9, 18
	stack.Push(array[len(array)-1])
	fmt.Println("<<", -1, ">>")

	for i := len(array) - 2; i >= 0; i-- {
		for !stack.IsEmpty() && stack.Top() <= array[i] {
			stack.Pop()
		}
		var res int
		if stack.IsEmpty() {
			res = -1
		} else {
			res = stack.Peek()
		}

		fmt.Println("<<", res, ">>")
		stack.Push(array[i])
	}
}

func NextGreaterElementCircular(array []int) []int {
	result := make([]int, 0)
	for range array {
		result = append(result, -1)
	}

	stack := make([]int, 0)

	for idx := 0; idx < 2*len(array); idx++ {
		circularIdx := idx % len(array)

		for len(stack) > 0 && array[stack[len(stack)-1]] < array[circularIdx] {
			var top int
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result[top] = array[circularIdx]
		}

		stack = append(stack, circularIdx)
	}

	return result
}

func main() {
	fmt.Println("Next Greatest element in the INDEX...")
	NextGreaterElementNaive([]int{5, 15, 10, 8, 6, 12, 9, 18})
}
