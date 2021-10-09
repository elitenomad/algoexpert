package main

import "fmt"

func NaivePatternSearch(s string, pattern string) {
	r := []rune(s)
	loops := len(pattern)
	for i := 0; i < len(s)-len(pattern)+1; i++ { //O(Length of string - Length of Pattern + 1) * M
		if string(r[i:i+loops]) == pattern {
			fmt.Printf("%d \t", i)
		}
	}
	fmt.Println()
}

func main() {
	a := "ABABABCD"
	p := "ABA"
	NaivePatternSearch(a, p)
	a = "ABCABCD"
	p = "ABCD"
	NaivePatternSearch(a, p)
	a = "AAAAA"
	p = "AAA"
	NaivePatternSearch(a, p)
}
