package main

import (
	"fmt"
	"strings"
)

// We can convert this into simple pattern matching algorithm
// And use simple string methods which are linear in time
func VerifyIfStrIsRotationOfAnother(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(s1+s1, s2)
}

func VerifyRotation(s1, s2 string) bool {
	r := []rune(s1)
	j := len(s1)

	for j > 0 { // O(N)
		first := r[0]

		for i := 1; i < len(r); i++ { // O(N)
			r[i-1] = r[i]
		}

		r[len(r)-1] = first

		if string(r) == s2 { // Comparison is O(1)
			return true
		}

		// Left rotate until you match
		j--
	}

	// if it doesnt match in all the rotations possible
	// which is length of the string we return false.
	return false
}

func main() {
	s1 := "ABCD"
	s2 := "CDAB"
	fmt.Println(VerifyRotation(s1, s2))
	// s1 = "ABAAA"
	// s2 = "BAAAA"
	// fmt.Println(VerifyRotation(s1, s2))
	// s1 = "ABAB"
	// s2 = "ABBA"
	// fmt.Println(VerifyRotation(s1, s2))
}
