package main

import "fmt"

// Ensure you print permutations of a strings
// in which a substring "AB" is not present

// Naive algorithm is
func Permutation(str string, left, right int) {
	r := []rune(str)

	if left == right {
		fmt.Println(str)
		return
	} else {
		for i := left; i <= right; i++ {
			r[left], r[i] = r[i], r[left]
			Permutation(string(r), left+1, right)
			r[left], r[i] = r[i], r[left]
		}
	}
}

func main() {
	a := "ABC"
	Permutation(a, 0, len(a)-1)
}
