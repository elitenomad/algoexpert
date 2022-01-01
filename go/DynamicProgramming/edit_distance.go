package main

import "fmt"

func EditDistance(s1, s2 string) int {
	return EditDistanceHelper(s1, s2, len(s1)-1, len(s2)-1)
}

func EditDistanceHelper(s1 string, s2 string, m, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if s1[m] == s2[n] {
		return EditDistanceHelper(s1, s2, m-1, n-1)
	} else {
		return 1 + min3(
			EditDistanceHelper(s1, s2, m, n-1),
			EditDistanceHelper(s1, s2, m-1, n),
			EditDistanceHelper(s1, s2, m-1, n-1),
		)
	}
}

func EditDistanceDP(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	dp[0][0] = 0
	for i := 1; i < len(dp); i++ {
		dp[0][i] = dp[0][i-1] + 1
	}

	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	fmt.Println(dp)
	return dp[len(s1)][len(s2)]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func main() {
	fmt.Println(EditDistanceDP("cat", "cut"))
	// fmt.Println(EditDistance("saturday", "sunday"))
}
