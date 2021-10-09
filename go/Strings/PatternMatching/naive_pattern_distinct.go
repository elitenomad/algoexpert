package main

import "fmt"

func NaivePatternDistinct(s string, pattern string) {
	for i := 0; i < len(s)-len(pattern)+1; {
		j := 0
		for j = 0; j < len(pattern); j++ {
			if pattern[j] != s[i+j] {
				break
			}
		}

		if j == len(pattern) {
			fmt.Printf("%d \t", i)
		}
		if j == 0 {
			i = i + 1
		} else {
			i = i + j
		}
	}
	fmt.Println()
}

func main() {
	a := "ABABABCD"
	p := "ABCD"
	NaivePatternDistinct(a, p)
	a = "ABCABCD"
	p = "ABCD"
	NaivePatternDistinct(a, p)
	a = "CDEFGHEFG"
	p = "EFG"
	NaivePatternDistinct(a, p)
}
