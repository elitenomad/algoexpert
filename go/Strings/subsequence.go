package main

import "fmt"

func PrintAllSubsequencesOfString(s string, output string) {
	if len(s) == 0 {
		fmt.Println(output)
		return
	}

	PrintAllSubsequencesOfString(s[1:], output+string(s[0]))
	PrintAllSubsequencesOfString(s[1:], output)
}

// O(N)
func IsS2SubsequenceOfS1Eff(s1, s2 string) bool {
	if s2 == "" {
		// Edge case when s2 is an empty string
		// For any string we have an emptry string has a subsequence
		return true
	}

	// Edge case where s1 length is less than s2
	if len(s1) < len(s2) {
		return false
	}

	i := 0
	j := 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
			j++
		} else {
			i++
		}
	}

	return (i == len(s1)) && (j == len(s2))
}

// Recursive exercises
func IsS2SubsequenceOfS1Recursive(s1 string, s2 string, n, m int) bool {
	if m == 0 {
		return true
	}

	if n == 0 {
		return false
	}

	if s1[n-1] == s2[m-1] {
		return IsS2SubsequenceOfS1Recursive(s1, s2, n-1, m-1)
	} else {
		return IsS2SubsequenceOfS1Recursive(s1, s2, n-1, m)
	}
}

func main() {
	a := "ABCD"
	b := "AD"
	output := ""
	PrintAllSubsequencesOfString(a, output)
	fmt.Println(IsS2SubsequenceOfS1Eff(a, b))
	fmt.Println(IsS2SubsequenceOfS1Recursive(a, b, len(a), len(b)))

}
