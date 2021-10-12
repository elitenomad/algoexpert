package main

import "fmt"

func ConvertZeroesToFives(input int) int {
	result := 0

	i := 0
	for input > 0 { //O(K)
		rem := input % 10
		if rem == 0 {
			rem = 5 // O(1)
		}

		result = rem*pow(10, i) + result
		i++
		input = input / 10
	}

	return result
}

func pow(val, n int) int {
	result := 1

	for n > 0 {
		result *= val
		n--
	}

	return result
}

func main() {
	a := 1004
	fmt.Println(ConvertZeroesToFives(a)) // 1554
	a = 121
	fmt.Println(ConvertZeroesToFives(a)) // 121
}
