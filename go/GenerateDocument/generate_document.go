package main

import "fmt"

func GenerateDocument(characters string, document string) bool {
	c := map[rune]int{}

	for _, cs := range characters {
		c[cs] += 1
	}

	fmt.Println(c)

	for _, cs := range document {
		if c[cs] == 0 {
			return false
		}

		c[cs] -= 1
	}
	fmt.Println(c)

	return true
}

func main() {
	c := "Bste!hetsi ogEAxpelrt x "
	b := "AlgoExpert is the Best!"

	fmt.Println(GenerateDocument(c, b))
}
