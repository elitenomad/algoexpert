package main

import "fmt"

/*
	Array Size Will always be n >= 2
	Only one element repeats in an array (any number of times)
	Only the elements from 0 to MAX(arr) are present
	0 <= MAX(arr) <= n - 2

	e.g.
	a := []int{0, 2, 1, 3, 2, 2}
	fmt.Println(RepeatingElement(a))
	3 possible solutions
	- Time O(N) Space O(N) (Using Array or Hash)
	- Time O(LOG(N)) Space O(1) (Sort and Loop through the array)
	- Time O(N) a	nd Space O(1) --- Head is exploding, Gonna look into solution tomorrow.
*/
func RepeatingElement(input []int) int { // 1 based indexing

	for i := 0; i < len(input); i++ {
		input[i], input[input[i]] = input[input[i]], input[i]

	}

	return -1
}

func repeatingElementInArray(arr []int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
}

func main() {
	a := []int{0, 2, 1, 3, 2, 2}
	fmt.Println(RepeatingElement(a))
}
