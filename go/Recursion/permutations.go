package main

import (
	"fmt"
)

func Permutation(s string) {
	PermutationHelper(s, 0)
}

func PermutationHelper(s string, i int) {
	if i == len(s)-1 {
		fmt.Println(s)
		return
	}

	for j := i; j < len(s); j++ {
		s = swap(s, i, j)
		PermutationHelper(s, i+1)
		s = swap(s, j, i)
	}
}

func swap(s string, i, j int) string {
	var r = []rune(s)

	temp := r[i]
	r[i] = r[j]
	r[j] = temp

	return string(r)
}

func main() {
	// fmt.Println(swap("ABCD", 0, 1))
	PermutationHelper("ABCD", 0)
}
