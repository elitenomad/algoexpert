package main

import "fmt"

func toh(n int, a, b, c string) {
	if n == 1 {
		fmt.Printf("Move 1 from %s to %s \n", a, c)
		return
	}

	toh(n-1, "A", "C", "B")
	fmt.Printf("Move %d from %s to %s \n", n, a, c)
	toh(n-1, "B", "A", "C")
}

func main() {
	toh(9, "A", "B", "C")
}
