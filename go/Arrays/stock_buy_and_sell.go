package main

import "fmt"

func StockBuyAndSell(input []int) int {
	maxProfit := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			maxProfit += input[i] - input[i-1]
		}
	}

	return maxProfit
}

func main() {
	a := []int{1, 5, 3, 8, 12}
	fmt.Println(StockBuyAndSell(a))
}
