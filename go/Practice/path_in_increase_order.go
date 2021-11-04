package main

import "fmt"

func PathInIncreasingOrder(N int, A []int, B []int) bool {
	if len(A) == 0 && len(B) == 0 {
		return false
	}
	h := map[string]bool{}
	for i, j := 0, 0; i < len(A); i, j = i+1, j+1 {
		min := min(A[i], B[j])
		max := max(A[i], B[j])
		h[fmt.Sprintf("%d", min)+fmt.Sprintf("%d", max)] = true
	}

	for i := 1; i < N; i++ {
		temp := fmt.Sprintf("%d", i) + fmt.Sprintf("%d", i+1)
		fmt.Println(temp)
		if _, found := h[temp]; !found {
			return false
		}
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	a := []int{1, 2, 4, 4, 3}
	b := []int{2, 3, 1, 3, 1}
	fmt.Println(PathInIncreasingOrder(4, a, b))
	a = []int{1, 2, 1, 3}
	b = []int{2, 4, 3, 4}
	fmt.Println(PathInIncreasingOrder(4, a, b))
	a = []int{2, 4, 5, 3}
	b = []int{3, 5, 6, 4}
	fmt.Println(PathInIncreasingOrder(6, a, b))
	a = []int{1, 3}
	b = []int{2, 2}
	fmt.Println(PathInIncreasingOrder(3, a, b))
	a = []int{}
	b = []int{}
	fmt.Println(PathInIncreasingOrder(2, a, b))
}
