package main

import "fmt"

func EditDistance(s1, s2 string) int {
	return EditDistanceHelper(s1, s2, len(s1)-1, len(s2)-1)
}

func EditDistanceHelper(s1 string, s2 string, m, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if s1[m] == s2[n] {
		return EditDistanceHelper(s1, s2, m-1, n-1)
	} else {
		return 1 + min3(
			EditDistanceHelper(s1, s2, m, n-1),
			EditDistanceHelper(s1, s2, m-1, n),
			EditDistanceHelper(s1, s2, m-1, n-1),
		)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func main() {
	fmt.Println(EditDistance("cat", "cut"))
	fmt.Println(EditDistance("saturday", "sunday"))
}
