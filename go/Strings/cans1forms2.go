package main

import "fmt"

/*
	Goal is to understand if we can form s2 using the number of chars in s1
*/
func CanS1FormS2(s1, s2 string) bool { // TC : O(S1 + S2), SC: O(S1 + S2)
	h1 := map[byte]int{}
	h2 := map[byte]int{}

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

func main() {
	a := "PRANAVA"
	b := "AVA"
	fmt.Println(IsS2SubsequenceOfS1(a, b))
}
