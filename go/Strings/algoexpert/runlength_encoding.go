package main

import (
	"fmt"
	"strconv"
)

func RunLengthEncoding(str string) string {
	// This method cannot be implemented because
	// Keys in the hash are not ordered and the results
	// Will be jumbled :)
	h := map[rune]int{}

	for _, char := range str {
		h[char] += 1
	}

	s := ""
	for k, v := range h {
		fmt.Println(string(k))
		q := v / 9
		r := v % 9

		if q > 0 {
			s = s + fmt.Sprint(9*q) + string(k) // append(runes, rune(1*q)+rune(k))
		}

		if r > 0 {
			s = s + fmt.Sprint(1*r) + string(k)
		}
	}

	return s
}

func RunLengthEncodingII(str string) string {
	encoders := []byte{}

	currentRunLength := 1
	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		previousChar := str[i-1]

		if (currentChar != previousChar) || currentRunLength == 9 {
			encoders = append(encoders, strconv.Itoa(currentRunLength)[0])
			encoders = append(encoders, previousChar)
			currentRunLength = 0
		}

		currentRunLength += 1
	}

	encoders = append(encoders, strconv.Itoa(currentRunLength)[0])
	encoders = append(encoders, str[len(str)-1])
	return string(encoders)
}

func main() {
	str := "AAAAAAAAAAAABBCCCCDD"
	fmt.Println(RunLengthEncoding(str))
	fmt.Println(RunLengthEncodingII(str))
}
