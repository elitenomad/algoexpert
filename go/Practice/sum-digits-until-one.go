// This is my solution for AlgoDaily problem Sum Digits Until One
// Located at https://algodaily.com/challenges/sum-digits-until-one

package main

import "fmt"

func sumDigitsUntilOne(arg int) int {
	result := arg
	sum := 0

	for result > 9 {
		temp := result

		for temp > 0 {
			r := temp % 10
			sum += r
			temp = temp / 10
		}

		result = sum
		sum = 0
	}

	return result
}

func main() {
	// write test cases
	fmt.Println(sumDigitsUntilOne(5000000))
}
