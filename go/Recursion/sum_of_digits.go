package main

import "fmt"

func sod(n int) int { // O(No of digits in the number) && space complexity is same as well.
	if n == 0 {
		return 0
	}

	temp := n % 10
	return temp + sod(n/10)
}

func main() {
	fmt.Println(sod(99876565665))
}
