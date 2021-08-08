package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func StockBuyAndSellBruteForce(input []int) int {
	start := 0
	end := len(input) - 1
	return StockBuyAndSellBruteForceHelper(input, start, end)
}

func StockBuyAndSellBruteForceHelper(input []int, start, end int) int {
	profit := 0

	for i := start; i < end; i++ {
		for j := i + 1; j <= end; j++ {
			if input[j] > input[i] {
				current_profit := input[j] - input[i] + StockBuyAndSellBruteForceHelper(input, start, i-1) + StockBuyAndSellBruteForceHelper(input, j+1, end)
				profit = max(profit, current_profit)
			}
		}
	}

	return profit
}

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
	fmt.Println(StockBuyAndSellBruteForce(a))
}
