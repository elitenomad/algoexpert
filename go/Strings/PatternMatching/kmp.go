package main

import "fmt"

// LPS is pre-requisite to understanding KMP algorithm

// Proper prefix
// "abc" - "" "a" "ab"
// "abc" - "" "c" "bc" "abc"
func LongestProperPrefixSuffixArray(s string, n int) int {
	for i := n - 1; i > 0; i-- {
		flag := true
		for j := 0; j < i; j++ {
			if s[i] != s[n-i+j] {
				flag = false
			}
		}

		if flag == true {
			return i
		}
	}

	return 0
}

func FillLPS(s string, lps []int) []int {
	for i := 0; i < len(s); i++ {
		lps[i] = LongestProperPrefixSuffixArray(s, i+1)
	}
	return lps
}

func main() {
	s := "abacabad"
	lps := make([]int, len(s))
	fmt.Println(FillLPS(s, lps))
}
