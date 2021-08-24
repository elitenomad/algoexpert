package main

import "fmt"

func BinaryNumber(n int) {
	for n > 0 {
		r := n % 2
		fmt.Println(r)
		n = n / 2
	}
}

func BinaryNumberRec(n int) {
	if n > 1 {
		BinaryNumberRec(n / 2)
	}
	fmt.Println(n % 2)
}

func main() {
	BinaryNumberRec(10)
}
