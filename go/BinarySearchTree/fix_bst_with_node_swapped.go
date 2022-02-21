package main

import "math"

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

/*
	Given an array with two moles resulting in unsorted array
	Find the values which are not in place, and put them in the right place
	4, 60, 10, 20, 8, 80, 100
*/

func FixBSTWithNodeSwappedValuesInArr(input []int) []int {
	prev := math.MinInt32
	first := -1
	second := -1

	for i := 0; i < len(input); i++ {
		if input[i] < prev {
			if first == -1 {
				first = prev
			}
			second = input[i]
		}
		prev = input[i]
	}

	return input
}

func FixBSTWithNodeSwapped(tree *BST) *BST {
	if tree == nil {
		return nil
	}

	FixBSTWithNodeSwapped(tree.Left)
	if Prev != nil && tree.Value < Prev.Value {
		if first == nil {
			first = Prev
		}
		second = tree
	}
	Prev = tree
	FixBSTWithNodeSwapped(tree.Right)
}

func main() {

}
