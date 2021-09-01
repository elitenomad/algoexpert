package main

import "fmt"

func contains(ints []int, val int) bool {
	for _, v := range ints {
		if v == val {
			return true
		}
	}

	return false
}

// Intersection of 2 unsorted arrays
// Use Hashing to solve the problem
func Intersection(a, b []int) int { // O(N) and O(N) for time and space complexities
	h := map[int]int{}
	count := 0

	for i := 0; i < len(a); i++ {
		h[a[i]] += 1
	}

	for j := 0; j < len(b); j++ {
		h[b[j]] += 1
	}

	for _, v := range h {
		if v > 1 {
			count += 1
		}
	}

	return count
}

func main() {
	a := []int{10, 15, 20, 5, 30}
	b := []int{30, 5, 30, 80}
	fmt.Println(Intersection(a, b))
	a = []int{10, 20}
	b = []int{20, 30}
	fmt.Println(Intersection(a, b))
	a = []int{10, 10, 10}
	b = []int{10, 10, 10}
	fmt.Println(Intersection(a, b))
}
