package main

import (
	"fmt"
	"math"
)

// You want to maximize your profit by choosing a single day to
// buy one stock and choosing a different day in the
// future to sell that stock.
func BestTimeToBuyAndSellStock(prices []int) int {
	minBuyPrice := math.MaxInt32
	maxProfit := math.MinInt32

	for i := 0; i < len(prices); i++ {
		minBuyPrice = min(minBuyPrice, prices[i])
		maxProfit = max(maxProfit, prices[i]-minBuyPrice)
	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(BestTimeToBuyAndSellStock(a))
}
