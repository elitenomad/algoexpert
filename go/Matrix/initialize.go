package main

import "fmt"

func Initialize2DArray() {
	one := [][]int{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 3},
	}
	fmt.Println(one)

	two := [3][2]int{}
	fmt.Println(len(two))

	three := [3][]int{}
	for i := 0; i < len(three); i++ {
		three[i] = make([]int, i+1)
		for j := 0; j < len(three[i]); j++ {
			three[i][j] = 2 * i
		}
		fmt.Println(three[i])
	}
	fmt.Println(three)
}

func main() {
	Initialize2DArray()
}
