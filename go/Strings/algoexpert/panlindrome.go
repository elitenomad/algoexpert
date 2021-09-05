package main

import "fmt"

/*
	Approach-1:
		Create an array of string length
		from last index append to the array created
		Join the array with emptry string
		compare two strings.
*/

func Palindrome(str string) bool { // O(N) Time and O(1) Space
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	str := "abcdcba"
	fmt.Println(Palindrome(str))
	str = "pranava"
	fmt.Println(Palindrome(str))
}
