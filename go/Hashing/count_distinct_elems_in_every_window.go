package main

import "fmt"

func CountDistinctElemsInWindow(input []int, k int) {
	// noOfWindows := len(input) - k + 1 //n - k + 1
	for i := 0; i < len(input)-k+1; i++ {
		h := map[int]bool{}
		for j := i; j < i+k; j++ {
			h[input[j]] = true
		}

		fmt.Println(len(h))
	}
}

func Efficient(input []int, k int) {
	h := map[int]int{}
	// Find the counts of first window
	for i := 0; i < k; i++ {
		h[input[i]] += 1
	}

	// Print first window's distinct elements
	fmt.Print("", len(h))

	for i := k; i < len(input); i++ {
		h[input[i-k]] -= 1
		if h[input[i-k]] == 0 {
			delete(h, input[i-k])
		}
		h[input[i]] += 1
		fmt.Print(" ", len(h))
	}
}

func main() {
	a := []int{10, 20, 20, 10, 30, 40, 10}
	CountDistinctElemsInWindow(a, 4)
	CountDistinctElemsInWindow(a, 7)
	fmt.Println("***")
	a = []int{10, 10, 10}
	CountDistinctElemsInWindow(a, 3)
	fmt.Println("***")
	a = []int{10, 20, 30, 40}
	CountDistinctElemsInWindow(a, 3)
	a = []int{10, 20, 10, 10, 30, 40}
	Efficient(a, 4)
}
