package main

import "fmt"

func Solution(S string) string {
	var occurrences [26]int
	for _, ch := range S {
		occurrences[int(ch)-int('a')]++
	}
	fmt.Println(occurrences)
	var best_char uint8 = 'a'
	var best_res int = 0
	var i int

	for i = 0; i < 26; i++ {
		if occurrences[i] > best_res {
			fmt.Println(uint8(int('a')+i), best_char)
			best_char = uint8(int('a') + i)
			best_res = occurrences[i]
		}
	}

	return string(best_char)
}

func main() {
	str := "b"
	fmt.Println(Solution(str))
}
