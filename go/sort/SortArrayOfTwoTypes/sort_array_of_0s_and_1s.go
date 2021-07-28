package main

import "fmt"

func Sort(input []int) []int {
	i := -1
	j := len(input)

	for {
		for ok := true; ok; ok = input[i] == 0 {
			i++
		}

		for ok := true; ok; ok = (input[j] == 1) {
			j--
		}

		if i >= j {
			break
		}

		input[i], input[j] = input[j], input[i]
	}

	return input
}

func main() {
	a := []int{0, 1, 1, 0, 1, 0, 1, 1}
	fmt.Println(Sort(a))
}
