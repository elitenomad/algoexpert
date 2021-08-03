package main

import (
	"fmt"
	"math"
)

func PowerSet(s string) {
	n := int(math.Pow(float64(2), float64(len(s))))

	for i := 0; i < n; i++ {
		for j := 0; j < len(s); j++ {
			if i&(1<<j) != 0 {
				fmt.Printf("%s", string(s[j]))
			}
		}
		fmt.Println()
	}
}

func main() {
	s := "abc"
	PowerSet(s)
}
