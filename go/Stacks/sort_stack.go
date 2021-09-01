package main

import "fmt"

/*
	Idea is to assume a stack as an ARRAY

	For e.g [-5, 2, -2, 4, 3, 1] - LIFO (Last IN First OUT)
	|	 1	|
	|	 3	|
	|	 4	|
	|	-2	|
	|	 2 	|
	|	-5	|
	Two Operation allowed : POP and PUSH
	In Place sort the array using stack operation

	PsuedoCode:
	POP and Element
	SORT Rest of the array of elements (Recursive)
	INSERT element to stack by value
		- If element stack is 0 or last element is <= value passed
	ELSE POP the element
	MODIFY THE EXISTING STACK
	Insert the Value to the stack
	Update the stack with latest top value
*/

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}

	top := stack[len(stack)-1]
	t_in := stack[:len(stack)-1]
	SortStack(t_in)

	InsertInSortedOrder(&stack, top)
	return stack
}

func InsertInSortedOrder(stack *[]int, value int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= value {
		*stack = append(*stack, value)
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	InsertInSortedOrder(stack, value)
	*stack = append(*stack, top)
}

func main() {
	fmt.Println("SortStack Implementation...")
	in := []int{-5, 2, -2, 4, 3, 1}
	out := SortStack(in)
	fmt.Println("SortStack Output: ", out)
}
