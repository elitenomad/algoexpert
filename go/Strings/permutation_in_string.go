package main

import (
	"fmt"
	"strings"
)

// Given two strings s1 and s2, return true if s2 contains a permutation of s1,
// or false otherwise. In other words, return true if one of s1's permutations
// is the substring of s2.

func checkInclusion(s1 string, s2 string) bool { // Solution exceeds TLE
	return checkInclusionHelper(s1, s2, 0, len(s1)-1)
}

func checkInclusionHelper(s string, t string, left, right int) bool {
	r := []rune(s)

	if left == right {
		return strings.Contains(t, string(r))
	} else {
		for i := left; i <= right; i++ {
			r[left], r[i] = r[i], r[left]
			value := checkInclusionHelper(string(r), t, left+1, right)
			if value == true {
				return true
			}
			r[left], r[i] = r[i], r[left]
		}
	}

	return false
}

// Second approach by taking the counts of char in both the strings

func checkInclusionII(s1 string, s2 string) bool {
	// Use case when s1 length > s2 length
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	if len(s1) > len(s2) {
		return false
	}

	s1Map := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]-'a'] += 1
	}
	fmt.Println(s1Map)

	for i := 0; i <= len(s2); i++ {
		s2Map := make([]int, 26)
		for j := 0; j < len(s1); j++ {
			s2Map[s2[i+j]-'a'] += 1
		}

		fmt.Println(s2Map)
		if matches(s1Map, s2Map) {
			return true
		}
	}

	return false
}

func matches(s1Map []int, s2Map []int) bool {
	for i := 0; i < 26; i++ {
		if s1Map[i] != s2Map[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	fmt.Println(checkInclusionII(s1, s2))
}
