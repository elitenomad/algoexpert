package main

import "fmt"

func Josephus(n, k int) int {
	return helper(n, k) + 1
}

func helper(n, k int) int {
	if n == 0 {
		return 0
	} else {
		return (helper(n-1, k) + k) % n
	}
}

func main() {
	fmt.Println(Josephus(5, 3))
}
