package main

import "fmt"

func Sort(input []int) []int {
	size := len(input)
	i := -1
	j := size

	for {
		for ok := true; ok; ok = (input[i]%2 == 0) {
			i++
		}

		for ok := true; ok; ok = !(input[j]%2 == 0) {
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
	a := []int{2, 3, 4, 5, 6}
	fmt.Println(Sort(a))
}
