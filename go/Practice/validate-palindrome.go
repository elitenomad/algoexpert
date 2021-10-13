// This is my solution for AlgoDaily problem Validate Palindrome
// Located at https://algodaily.com/challenges/validate-palindrome

package main

import "fmt"

func validatePalindrome(arg string) bool {
	for i, j := 0, len(arg)-1; i < len(arg)/2; i, j = i+1, j-1 {
		if arg[i] != arg[j] {
			return false
		}
	}
	return true
}

func main() {
	// write test cases
	fmt.Println(validatePalindrome("racecar"))
	fmt.Println(validatePalindrome("thisisnotapalindrome"))

}
