package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string { //
	splitted := strings.Split(s, " ") // O(N)

	for i, j := 0, len(splitted)-1; i < len(splitted)/2; i, j = i+1, j-1 { // O(N/2)
		splitted[i], splitted[j] = splitted[j], splitted[i]
	}

	return strings.Join(splitted, " ") // O(N)
}

// Another approach is to use stack
// Parse through the charectors when whitespace is encountered means a word
// Push the word to the stack
// Once you reach the end of the string, exit
// Pop each word from the stack and append it to orignal string

func ReverseWordsII(str string) string {
	words := []string{}
	start := 0

	for i, char := range str { // O(N)
		if char == ' ' {
			words = append(words, str[start:i])
			start = i
		}
	}

	words = append(words, str[start:])

	for i := 0; i < len(words); i++ { // O(Words)
		words[i] = reverse(words[i])
	}

	temp := strings.Join(words, " ")
	result := []byte{}
	for i := len(temp) - 1; i >= 0; i-- { // O(String)
		result = append(result, temp[i])
	}
	// O(Length of String) + O(Length of Words)
	return string(result)
}

func reverse(input string) string {
	r := []rune(input)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	a := "tim is great"
	fmt.Println(ReverseWordsII(a))
}
