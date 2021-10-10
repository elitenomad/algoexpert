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
func IsAnagram(s1 string, pattern string) bool {
	if len(s1) != len(pattern) {
		return false
	}

	chars := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		chars[s1[i]-'a'] += 1
		chars[pattern[i]-'a'] -= 1
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}
func VerifyIfAnagramIsPresent(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	for i := 0; i < n-m+1; i++ {
		if IsAnagram(s1[i:i+m], s2) {
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

func VerifyIfAnagramIsPresentII(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	s1Slice := make([]int, 256)
	s2Slice := make([]int, 256)

	for i := 0; i < m; i++ {
		s1Slice[s1[i]] += 1
		s2Slice[s2[i]] += 1
	}

	for i := m; i < n; i++ {

		if matches(s1Slice, s2Slice) {
			return true
		}

		s1Slice[s1[i]] += 1
		s1Slice[s1[i-m]] -= 1
	}

	return false
}

func main() {
	// s1 := "eidbaooo"
	// s2 := "ab"
	// fmt.Println(VerifyIfAnagramIsPresent(s1, s2))
	s1 := "geeksfargeeks"
	s2 := "frog"
	fmt.Println(VerifyIfAnagramIsPresentII(s1, s2))
}
