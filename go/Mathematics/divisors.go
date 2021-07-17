package main

import (
	"fmt"
)

/*
	Naive solution is to go from 1 to n
	And verify if the numer could be divisible
	by each number in the loop. Print the number
	if the remainder is 0
*/
func Divisors(a int) { // O(n) O(1)
	for i := 1; i < a; i++ {
		if a%i == 0 {
			fmt.Printf("%d \t", i)
		}
	}
}

func DivisorsOptimized(a int) { // O(SqrtN) + O(SqrtN) = O(SqrtN)
	i := 1
	for i = 1; i*i < a; i++ {
		if a%i == 0 {
			fmt.Printf("%d \t", i)
		}
	}
	fmt.Println("I: is :", i)
	for i >= 1 {
		if a%i == 0 {
			fmt.Printf("%d \t", a/i)
		}
		i -= 1
	}
}

func main() {
	a := 15
	DivisorsOptimized(a)
}
