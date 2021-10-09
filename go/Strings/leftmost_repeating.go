package main

import (
	"fmt"
	"sort"
)

// We cannot know until we loop throuhg
// All the chars of the string
type str [2]int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func LeftMostRepeatChar(s string) int {
	h := map[byte]str{} // Auxillary space O(N) where N is the length of the string

	for i := 0; i < len(s); i++ { // O(N)
		if _, found := h[s[i]]; found {
			v := h[s[i]] // Fetch first
			v[0] = min(v[0], i)
			v[1] += 1
			h[s[i]] = v
		} else {
			h[s[i]] = str{i, 1}
		}
	}

	// Now loop through the h with values > 1
	a := []int{}
	for _, v := range h { // O(N) - N is the size of the String
		if v[1] > 1 {
			a = append(a, v[0])
		}
	}

	sort.Ints(a) // N Log N

	if len(a) <= 0 {
		return -1
	}

	return a[0]
}

func LeftMostRepeatCharII(s string) int {
	h := make([]byte, 26)

	for i := 0; i < len(s); i++ {
		h[s[i]-'a'] += 1
	}

	for i := range h {
		if h[i] > 1 {
			return i
		}
	}

	return -1
}

func LeftMostRepeatCharIII(s string) int {
	input := make([]int, 26)
	for i := 0; i < len(input); i++ {
		input[i] = -1
	}

	res := 100000000
	for i := 0; i < len(s); i++ {
		if input[s[i]-'a'] == -1 {
			input[s[i]-'a'] = i
		} else {
			res = min(res, input[s[i]-'a'])
		}
	}

	if res == 100000000 {
		return -1
	}
	return res
}

// Another efficient apprach is
// Traverse the string from right to left
// maintain visited array
// create a variable result
// Update the result when you see which is already visited
// In the end we return the result

func main() {
	a := "abbcc"
	fmt.Println(LeftMostRepeatChar(a))
	fmt.Println(LeftMostRepeatCharIII(a))
	a = "geeksforgeeks"
	fmt.Println(LeftMostRepeatChar(a))
	fmt.Println(LeftMostRepeatCharIII(a))
	a = "abcd"
	fmt.Println(LeftMostRepeatChar(a))
	fmt.Println(LeftMostRepeatCharIII(a))
	a = "abccb"
	fmt.Println(LeftMostRepeatChar(a))
	fmt.Println(LeftMostRepeatCharIII(a))
}
