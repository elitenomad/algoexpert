package main

import "fmt"

func generate_subsets(s string) {
	helper(s, "", 0)
}

func helper(s string, current string, i int) {
	if i == len(s) {
		fmt.Println(current)
		return
	}

	helper(s, current, i+1)
	helper(s, current+string(s[i]), i+1)
}

func GenerateSubsetsGivenArray(input []int, size int) [][]int {
	result := [][]int{}
	temp := []int{}
	GenerateSubsetsGivenArrayHelper(input, temp, &result, 0)
	return result
}

func GenerateSubsetsGivenArrayHelper(input []int, temp []int, result *[][]int, size int) {
	if size == len(input) {
		*result = append(*result, temp)
		return
	}

	GenerateSubsetsGivenArrayHelper(input, temp, result, size+1)
	temp = append(temp, input[size])
	GenerateSubsetsGivenArrayHelper(input, temp, result, size+1)

}

func main() {
	// generate_subsets("ABCD")
	fmt.Println(GenerateSubsetsGivenArray([]int{2, 5, 10}, 0))
}
