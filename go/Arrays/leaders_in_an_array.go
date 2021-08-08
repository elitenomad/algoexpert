package main

import "fmt"

/*
	Brute force approach
*/
func LeadersInArray(input []int) []int { // O(N^2)
	temp := []int{}

	for i := 0; i < len(input); i++ {
		flag := true // Assume all are leaders

		for j := i + 1; j < len(input); j++ {
			if input[i] <= input[j] {
				flag = false
				break
			}
		}

		if flag {
			temp = append(temp, input[i])
		}

		flag = true
	}

	return temp
}

// Only thing is this solution displays the elements
// From right to left
// We can sort in the end if we need elements from right to left
func LeadersInArrayEff(input []int) []int { //O(n)
	lastIndex := len(input) - 1
	temp := []int{}

	currentLeader := input[lastIndex]
	temp = append(temp, currentLeader)

	for i := lastIndex - 1; i >= 0; i-- {
		if currentLeader < input[i] {
			currentLeader = input[i]
			temp = append(temp, currentLeader)
		}
	}

	return temp
}

func main() {
	a := []int{7, 10, 4, 10, 6, 5, 2}
	fmt.Println(LeadersInArray(a))
	fmt.Println(LeadersInArrayEff(a))
}
