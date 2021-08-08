package main

import "fmt"

func Frequencies(input []int) map[int]int {
	h := map[int]int{}

	for i := 0; i < len(input); i++ {
		if h[input[i]] > 0 {
			h[input[i]] += 1
		} else {
			h[input[i]] = 1
		}
	}

	return h
}

/*
	Another approach without using the map would be just have a frequency variable
	Print the value of frequency if the current element != previous element
	Reset the value of the frequency to 1 and start over if the current_element != previous element
	Repeat the 2 and 3 until the looping of array is complete
*/

func main() {
	a := []int{10, 10, 10, 25, 30, 30}
	fmt.Println(Frequencies(a))
	a = []int{40, 50, 50, 50}
	fmt.Println(Frequencies(a))
	a = []int{10}
	fmt.Println(Frequencies(a))
}
