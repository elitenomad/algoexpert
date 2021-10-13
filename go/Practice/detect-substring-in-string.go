// This is my solution for AlgoDaily problem Detect Substring in String
// Located at https://algodaily.com/challenges/detect-substring-in-string

package main

import "fmt"

func detectSubstringInString(str, substr string) int {
	for i := 0; i < len(str); i++ {
		if str[i] == substr[0] {
			if str[i:i+len(substr)] == substr {
				return i
			}
		}
	}
	return -1
}

func main() {
	// write test cases
	fmt.Println(detectSubstringInString("carnival", "jiiva"))
}
