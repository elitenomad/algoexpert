package main

import "fmt"

// Using Recursion
func LongestCommonSubsequence(s1, s2 string) int {
	return LongestCommonSubsequenceHelper(s1, s2, len(s1), len(s2))
}

func LongestCommonSubsequenceHelper(s1, s2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if s1[m-1] == s2[n-1] {
		return 1 + LongestCommonSubsequenceHelper(s1, s2, m-1, n-1)
	} else {
		return max(
			LongestCommonSubsequenceHelper(s1, s2, m-1, n),
			LongestCommonSubsequenceHelper(s1, s2, m, n-1),
		)
	}
}

// Recursion + caching = Top Down approach
func LongestCommonSubsequenceII(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	return LongestCommonSubsequenceHelperII(s1, s2, len(s1)-1, len(s2)-1, &cache)
}

func LongestCommonSubsequenceHelperII(s1, s2 string, m, n int, cache *[][]int) int {
	if (*cache)[m][n] != -1 {
		return (*cache)[m][n]
	}

	if m == 0 || n == 0 {
		(*cache)[m][n] = 0
	} else {
		if s1[m-1] == s2[n-1] {
			(*cache)[m][n] = 1 + LongestCommonSubsequenceHelperII(s1, s2, m-1, n-1, cache)
		} else {
			(*cache)[m][n] = max(
				LongestCommonSubsequenceHelperII(s1, s2, m-1, n, cache),
				LongestCommonSubsequenceHelperII(s1, s2, m, n-1, cache),
			)
		}
	}

	return (*cache)[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func LongestCommonSubsequenceIII(s1, s2 string) int {
	m := len(s1)
	n := len(s2)

	result := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		result[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		result[i][0] = 0
	}

	for i := 0; i <= n; i++ {
		result[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				result[i][j] = 1 + result[i-1][j-1]
			} else {
				result[i][j] = max(result[i][j-1], result[i-1][j])
			}
		}
	}
	return result[m][n]
}

/*
	Diff Utility
	Minimum deletions and insertions to convert s1 to s2
	Shortest common supersequence
	Longest Palindrom subsequence
	Longest Repeating subsequence
	Space optimized dp solution for LCS
	Printing LCS
*/

func main() {
	s1 := "ABCDGH"
	s2 := "AEDFHR"
	fmt.Println(LongestCommonSubsequence(s1, s2))
	fmt.Println(LongestCommonSubsequenceII(s1, s2))
	fmt.Println(LongestCommonSubsequenceIII(s1, s2))
}
