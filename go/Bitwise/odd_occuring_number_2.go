package main

import "fmt"

func OddOccuringNumberTwice(input []int) []int {
	result := make([]int, 2)
	i := 0
	h := map[int]int{}

	// Loop through an array
	for i := 0; i < len(input); i++ {
		if _, ok := h[input[i]]; ok {
			h[input[i]] = h[input[i]] + 1
		} else {
			h[input[i]] = 1
		}
	}

	for key, val := range h {
		if val%2 != 0 {
			result[i] = key
			i++
		}
	}

	return result
}

func OddOccuringNumberUsingXOR(input []int) (int, int) {
	xor := 0
	res_1 := 0
	res_2 := 0

	for i := 0; i < len(input); i++ {
		xor = xor ^ input[i]
	}

	firstSetBit := xor & ^(xor - 1)

	for i := 0; i < len(input); i++ {
		if (input[i] & firstSetBit) != 0 {
			res_1 = res_1 ^ input[i]
		} else {
			res_2 = res_2 ^ input[i]
		}
	}

	return res_1, res_2
}

func main() {
	a := []int{3, 4, 3, 4, 5, 4, 4, 6, 7, 7}
	fmt.Println(OddOccuringNumberTwice(a))
	fmt.Println(OddOccuringNumberUsingXOR(a))
}
