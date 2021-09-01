package main

import "fmt"

// Assuming the array is increasing order
func IsArraySorted(input []int) bool {
	for i := 0; i < len(input)-1; i++ {
		if input[i] > input[i+1] {
			return false
		}
	}

	return true
}
func main() {
	a := []int{10, 20, 30}
	fmt.Println(IsArraySorted(a))
	a = []int{200, 20, 100}
	fmt.Println(IsArraySorted(a))
	a = []int{200, 100, 50}
	fmt.Println(IsArraySorted(a))
}
