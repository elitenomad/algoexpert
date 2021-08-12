package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	var store = make(map[int]bool)

	for _, elem := range array {
		sub := target - elem

		if _, found := store[sub]; found {
			return []int{sub, elem}
		}

		store[elem] = true
	}

	return []int{}
}

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	givenSum := 10
	fmt.Println(TwoNumberSum(a, givenSum))
}
