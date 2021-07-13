package main

import "fmt"

func FirstNonRepeatingCharacter(str string) int {
	c := map[rune]int{}
	for _, v := range str {
		c[v] += 1
	}

	for i, v := range str {
		if c[v] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	i_1 := "prranavan"
	fmt.Println(FirstNonRepeatingCharacter(i_1))
	i_2 := "aabbcc"
	fmt.Println(FirstNonRepeatingCharacter(i_2))
}
