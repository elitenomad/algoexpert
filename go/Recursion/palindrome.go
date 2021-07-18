package main

import (
	"fmt"
)

func IsPalindrome(s string) bool {
	return IsPalindromeHelper(s, 0, len(s)-1)
}

func IsPalindromeHelper(s string, begin, end int) bool { // O(n) -> T and O(n) -> S
	if begin >= end {
		return s[begin] == s[end] //
	}

	return (s[begin] == s[end]) && IsPalindromeHelper(s, begin+1, end-1)
}

func main() {
	fmt.Println(IsPalindrome("abba"))
	fmt.Println(IsPalindrome("abbcbba"))
	fmt.Println(IsPalindrome("geeks"))
}
