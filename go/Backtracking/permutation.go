package main

import "fmt"

// Ensure you print permutations of a strings
// in which a substring "AB" is not present

// Naive algorithm is
func Permutation(str string, left, right int) {
	r := []rune(str)

	if left == right {
		fmt.Println(str)
		return
	} else {
		for i := left; i <= right; i++ {
			r[left], r[i] = r[i], r[left]
			Permutation(string(r), left+1, right)
			r[left], r[i] = r[i], r[left]
		}
	}
}

func PermutationHelper(nums []int, ds []int, freq []bool, result *[][]int) {
	if len(ds) == len(nums) {
		(*result) = append((*result), ds)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(freq)[i] {
			(freq)[i] = true
			(ds) = append((ds), nums[i])
			PermutationHelper(nums, ds, freq, result)
			ds = (ds)[0 : len(ds)-1]
			(freq)[i] = false
		}
	}
}

func PermutationII(nums []int) [][]int {
	result := [][]int{}
	frequencies := make([]bool, len(nums))
	ds := []int{}

	// Start the recursive function
	PermutationHelper(nums, ds, frequencies, &result)

	return result
}

func main() {
	// a := "ABC"
	// Permutation(a, 0, len(a)-1)
	nums := []int{1, 2, 3, 4}
	fmt.Println(PermutationII(nums))
}
