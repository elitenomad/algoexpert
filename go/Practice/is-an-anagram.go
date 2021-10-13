// This is my solution for AlgoDaily problem Is An Anagram
// Located at https://algodaily.com/challenges/is-an-anagram

package main

import "fmt"

func isAnAnagram(s1, s2 string) bool {
	chars := make([]int, 256)

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		chars[s1[i]] += 1
		chars[s2[i]] -= 1
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnAnagram("cinema", "iceman"))
}
