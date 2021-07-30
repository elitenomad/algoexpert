package main

import (
	"fmt"
	"math/bits"
)

func Naive(n int) int {
	res := 0

	for n > 0 {
		if n%2 != 0 {
			res++
		}

		n = n / 2
	}

	return res
}

func NaiveBitWiseOperators(n int) int { // O(8) || O(16) || O(32) || O(64)
	res := 0

	for n > 0 {
		if n&1 == 1 {
			res++
		}

		n = n >> 1
	}

	return res
}

/*
	Brian kennigam's Algorithm.
	Turn off the last bit set in every
	Iteration
*/
func SetBitCount(n int) int {
	res := 0

	for n > 0 {
		n = (n & (n - 1))
		res++
	}

	return res
}

/*
	Lookup Table Algorithm
*/
var table [256]int

func initializeTable() {
	table[0] = 0
	for i := 1; i < 256; i++ {
		table[i] = table[i/2] + (i & 1)
	}
}

func LookupTableMethod(n int) int {
	initializeTable()
	var res int

	res = table[(n & 0xff)]
	n = n >> 8
	res = res + table[n&0xff]
	n = n >> 8
	res = res + table[n&0xff]
	n = n >> 8
	res = res + table[n&0xff]
	n = n >> 8

	return res
}

func main() {
	// In built Go Methods
	fmt.Printf("OnesCount(%b) = %d\n", 14, bits.OnesCount(14))
	fmt.Printf("OnesCount(%b) = %d\n", 5, bits.OnesCount(5))
	fmt.Printf("OnesCount(%b) = %d\n", 7, bits.OnesCount(7))

	// Naive Implementation
	fmt.Println("set Bits for: 7", Naive(7))

	// Brian kenningams Algorithm
	fmt.Println(SetBitCount(40))

	// Lookup Table Algorithm
	fmt.Println(LookupTableMethod(13)) // 3
}
