package main

import "fmt"

func KthBitSetForGivenNo(n, k int) bool {
	if (n & (1 << (k - 1))) != 0 { // (n >> k - 1) & 1 == 1
		return true
	}

	return false
}

func main() {
	n := 5
	k := 2
	fmt.Println(KthBitSetForGivenNo(n, k))
}
