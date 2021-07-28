package main

import "fmt"

/*
	Problem with below solution is that we loop
	through an array multiple times.
*/
func SortUsingNaiveMethod(input []int) []int {
	l := len(input)
	temp := make([]int, l)

	i := 0
	for j := 0; j < l; j++ {
		if input[j] < 0 {
			temp[i] = input[j]
			i++
		}
	}

	for j := 0; j < l; j++ {
		if input[j] > 0 {
			temp[i] = input[j]
			i++
		}
	}

	return temp
}

/*
	Use Lumoto or Hoares partition to solve the problem
*/
func Sort(input []int) []int {
	i := -1
	j := len(input)

	for {
		for ok := true; ok; ok = (input[i] < 0) {
			i++
		}

		for ok := true; ok; ok = (input[j] > 0) {
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
	a := []int{13, -12, 8, -10}
	fmt.Println(SortUsingNaiveMethod(a))
	fmt.Println(Sort(a))
}
