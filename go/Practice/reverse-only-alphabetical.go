// This is my solution for AlgoDaily problem Reverse Only Alphabetical
// Located at https://algodaily.com/challenges/reverse-only-alphabetical

package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyAlphabetical(arg string) string {
	output := make([]rune, len(arg))
	input := []rune(arg)

	if len(input) <= 0 {
		return ""
	}

	i := 0
	for j := len(input) - 1; j >= 0; j-- {
		if unicode.IsLetter(input[j]) {
			if unicode.IsLetter(input[i]) {
				output[i] = input[j]
				i += 1
			} else {
				i += 1
				output[i] = input[j]
				i += 1
			}
		} else {
			output[j] = input[j]
		}
	}
	return string(output)
}

func main() {
	// write test cases
	fmt.Println(reverseOnlyAlphabetical("B!FDCEA2"))
	fmt.Println(reverseOnlyAlphabetical("sea!$hells3"))
}
