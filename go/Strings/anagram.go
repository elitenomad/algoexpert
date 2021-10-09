package main

import "fmt"

// Sort two strings and Compare -- O(NLogN)
func Anagram(s1, s2 string) bool {
	h1 := map[byte]int{}
	h2 := map[byte]int{}

	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		h1[s1[i]] += 1
	}

	for i := 0; i < len(s2); i++ {
		h2[s2[i]] += 1
	}

	for k, v := range h2 {
		if _, found := h1[k]; !found {
			return false
		}

		if v > h1[k] {
			return false
		}
	}

	return true
}

/*
	Another approach to have constant Auxillary space
	Create a char array with int counts
	Loop through First array and then increament the char counts
	Loop through second array and then decrement the char counts
	Loop through CHar counts
		- If its not 0  then return false
		- Else return true
*/

func AnagramWithConstantSpace(s string, t string) bool {
	r := make([]rune, 26)

	for i := 0; i < len(s); i++ {
		r[s[i]-'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		r[t[i]-'a'] -= 1
	}

	for i := 0; i < len(r); i++ {
		if r[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	a := "silent"
	b := "listen"
	fmt.Println(Anagram(a, b))
	fmt.Println(AnagramWithConstantSpace(a, b))
}
