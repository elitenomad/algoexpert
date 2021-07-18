package main

import "fmt"

func generate_subsets(s string) {
	helper(s, "", 0)
}

func helper(s string, current string, i int) {
	if i == len(s) {
		fmt.Println(current)
		return
	}

	helper(s, current, i+1)
	helper(s, current+string(s[i]), i+1)
}

func main() {
	generate_subsets("ABCD")
}
