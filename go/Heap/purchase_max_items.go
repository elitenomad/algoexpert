package main

import (
	"fmt"
	"sort"
)

func PurchaseMaxItems(costs []int, sum int) int { // N Logn
	result := 0
	sort.Ints(costs)
	temp := sum
	fmt.Println(costs)
	for i := 0; i < len(costs); i++ {
		if temp > costs[i] {
			temp -= costs[i]
			result += 1
		} else {
			break
		}
	}

	return result
}

/*
	Another approach is convert the given array it into Binary Heap
	Turn it into Min Heap
	Extract Min until sum > popped element (while doing so result+=1)
	O(N) (To Build a Heap - Technicall its O(NLOGN)) + O(RES*LOGN)
*/

func main() {
	input := []int{1, 12, 5, 111, 200}
	fmt.Println(PurchaseMaxItems(input, 10))
}
