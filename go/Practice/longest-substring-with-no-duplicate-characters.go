// This is my solution for AlgoDaily problem Longest Substring With No Duplicate Characters
// Located at https://algodaily.com/challenges/longest-substring-with-no-duplicate-characters

package main

import "fmt"

type Substring struct {
	Left  int
	Right int
}

func (ss Substring) length() int {
	return ss.Right - ss.Left
}

func LongestSubstring(str string) string {
	substr := Substring{0, 1}
	h := map[rune]int{}
	start := 0

	for i, char := range str {
		if seenIndex, found := h[char]; found && start < seenIndex+1 {
			start = seenIndex + 1
		}

		if substr.length() < i+1-start {
			substr = Substring{start, i + 1}
		}

		h[char] = i
	}

	return str[substr.Left:substr.Right]
}

func main() {
	fmt.Println(LongestSubstring("abracadabar"))
}
