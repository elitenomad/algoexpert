package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	// Look for backspace
	s_rune := []rune(s)
	r_1 := ""
	for i := 0; i < len(s_rune)-1; i++ {
		if string(s_rune[0]) == "#" || string(s_rune[len(s_rune)-1]) == "#" {
			continue
		}

		if string(s_rune[i]) != "#" && string(s_rune[i+1]) != "#" {
			r_1 += string(s_rune[i])
		}
	}

	t_rune := []rune(t)
	r_2 := ""
	for i := 0; i < len(t_rune)-1; i++ {
		if string(t_rune[0]) == "#" || string(t_rune[len(t_rune)-1]) == "#" {
			continue
		}

		if string(t_rune[i]) != "#" && string(t_rune[i+1]) != "#" {
			r_2 += string(t_rune[i])
		}
	}

	return r_1 == r_2
}

func main() {
	a := "xywrrmp"
	b := "xywrrmu#p"
	fmt.Println(backspaceCompare(a, b))
}
