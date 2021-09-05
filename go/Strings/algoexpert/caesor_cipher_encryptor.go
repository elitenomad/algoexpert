package main

import (
	"fmt"
	"strings"
)

func Caesor(str string, k int) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	runes := []rune(str)
	fmt.Println(runes)

	for i, char := range runes {
		index := strings.Index(alphabets, string(char))
		if index == -1 {
			return ""
		}
		newIndex := (index + k) % 26
		runes[i] = rune(alphabets[newIndex])
	}

	return string(runes)
}

func main() {
	str := "xyz"
	fmt.Println(Caesor(str, 2))
}
