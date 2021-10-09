package main

import (
	"fmt"
)

func LeftMostNonRepeatChar(s string) int {
	h := make([]byte, 26)

	for i := 0; i < len(s); i++ {
		h[s[i]-'a'] += 1
	}

	for i := range h {
		if h[i] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	a := "abbcc"
	fmt.Println(LeftMostNonRepeatChar(a))
	a = "geeksforgeeks"
	fmt.Println(LeftMostNonRepeatChar(a))
	a = "abcd"
	fmt.Println(LeftMostNonRepeatChar(a))
	a = "aabccb"
	fmt.Println(LeftMostNonRepeatChar(a))
}
