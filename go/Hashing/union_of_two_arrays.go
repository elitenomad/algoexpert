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

// Union of 2 unsorted arrays
// Use Hashing to solve the problem
func Union(a, b []int) int { // O(N) and O(N) for time and space complexities
	h := map[int]bool{}

	for i := 0; i < len(a); i++ {
		h[a[i]] = true
	}

	for j := 0; j < len(b); j++ {
		h[b[j]] = true
	}

	return len(h)
}

func main() {
	a := []int{10, 15, 20, 5, 30}
	b := []int{30, 5, 30, 80}
	fmt.Println(Union(a, b))
	a = []int{10, 20}
	b = []int{20, 30}
	fmt.Println(Union(a, b))
	a = []int{10, 10, 10}
	b = []int{10, 10, 10}
	fmt.Println(Union(a, b))
}
