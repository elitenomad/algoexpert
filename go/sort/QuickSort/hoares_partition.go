package main

import "fmt"

func HoaresPartition(input []int) int {
	len := len(input)
	first := 0
	last := len - 1
	i := first - 1
	j := last + 1
	pivot := input[first]

	for {
		for ok := true; ok; ok = (input[i] < pivot) {
			i++
		}

		for ok := true; ok; ok = (input[j] > pivot) {
			j--
		}

		if i >= j {
			return j
		}
		input[i], input[j] = input[j], input[i]
	}
}

func main() {
	a := []int{3, 8, 6, 12, 10, 7}
	fmt.Println(HoaresPartition(a))
	a = []int{70, 60, 80, 40, 30}
	fmt.Println(HoaresPartition(a))
	a = []int{30, 40, 20, 50, 80}
	fmt.Println(HoaresPartition(a))
}
