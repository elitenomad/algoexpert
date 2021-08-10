package main

import "fmt"

func MinimumGroupFlips(input []int) {

	for i := 1; i < len(input); i++ {
		if input[i] != input[i-1] {
			if input[i] != input[0] {
				fmt.Print("From ", i, " to")
			} else {
				fmt.Println("", i-1)
			}
		}

		if input[len(input)-1] != input[0] {
			fmt.Println(len(input) - 1)
		}
	}
}

func main() {
	a := []int{0, 0, 1, 1, 0, 0, 1, 1, 0}
	MinimumGroupFlips(a)
}
