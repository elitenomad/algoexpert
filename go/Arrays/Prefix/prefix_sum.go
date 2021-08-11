package main

import "fmt"

/*
	Given an array how to perform the sum operations in the range
	efficiently

	e.g input = []int{2,8,3,9,6,5,4}
	getSum(0,2)
	getSum(1,3)
	getSum(4,6)
*/
func GetSum(input []int, l, r int) int {
	prefix := []int{}
	sum := 0
	result := 0

	for i := 0; i < len(input); i++ {
		sum += input[i]
		prefix = append(prefix, sum)
	}
	// fmt.Println(prefix) // {2, 10, 13, 22, 28, 33, 37}

	if l != 0 {
		result = prefix[r] - prefix[l-1]
	} else {
		result = prefix[r]
	}

	return result
}

func main() {
	a := []int{4, 8, 3, 9, 6, 5, 4}
	fmt.Println(GetSum(a, 0, 2))
	fmt.Println(GetSum(a, 2, 3))
	fmt.Println(GetSum(a, 4, 6))
}
