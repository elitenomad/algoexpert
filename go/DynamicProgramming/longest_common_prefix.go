package main

import (
	"fmt"
	"strings"
)

func LongestCommonPrefix(input []string) string {
	if len(input) == 0 {
		return ""
	}

	prefix := input[0]
	for i := 1; i < len(input); i++ {
		for strings.Index(input[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func LongestCommonPrefixII(input []string) string {
	return LongestCommonPrefixHelper(input, 0, len(input)-1)
}

func LongestCommonPrefixHelper(input []string, i, j int) string {
	if i == j {
		return input[i]
	} else {
		mid := (i + j) / 2
		b := LongestCommonPrefixHelper(input, i, mid)
		c := LongestCommonPrefixHelper(input, mid+1, j)
		return commonPrefix(b, c)
	}
}

func commonPrefix(s1, s2 string) string {
	m := min(len(s1), len(s2))

	for i := 0; i < m; i++ {
		if s1[i] != s2[i] {
			return s1[0:i]
		}
	}

	return s1[0:m]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func main() {
	strs := []string{"flow", "flower", "flight"}
	fmt.Println(LongestCommonPrefixII(strs))
}
