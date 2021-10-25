package main

// you can also use imports, for example:
// import "fmt"
import (
	"fmt"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func BinaryNumber(n int) []rune {
	runes := []rune{}
	for n > 0 {
		r := n % 2
		runes = append(runes, rune(r))
		n = n / 2
	}

	return runes
}

func Solution(N int) int {
	// write your code in Go 1.4
	nums := BinaryNumber(N)

	arr := []int{}
	result := 0

	for i := 0; i < len(nums); i++ {
		if int(nums[i]) == 1 {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
	if len(arr) == 1 {
		return 0
	}

	if len(arr) == 2 {
		return arr[1] - arr[0] - 1
	}

	for i := 1; i < len(arr); i++ {
		result = max(result, arr[i]-arr[i-1]-1)
	}
	return result
}

func binaryGap(N int) int {
	A := [32]int{}
	t := 0

	for i := 0; i < 32; i++ {
		if ((N >> i) & 1) != 0 {
			t += 1
			A[t] = i
		}
	}
	fmt.Println(t, A)

	ans := 0
	for i := 0; i < t-1; i++ {
		ans = max(ans, A[i+1]-A[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	a := 9
	fmt.Println(binaryGap(a))
}
