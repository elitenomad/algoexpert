package main

import "fmt"

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

func TrapRainWater(input []int) int {
	amount := 0
	// When its decreasing or when its increasing
	// Amount of water we trap is 0

	// Goal is to return the amount of water we trap
	// B/w the rectangular boxes
	for i := 1; i < len(input)-1; i++ {
		lmax := input[i]
		for j := 0; j < i; j++ {
			lmax = max(lmax, input[j])
		}

		rmax := input[i]
		for j := i + 1; j < len(input); j++ {
			rmax = max(rmax, input[j])
		}

		amount = amount + (min(lmax, rmax) - input[i])
	}

	return amount
}

func TrapRainWaterEff(input []int) int {
	// Calculate lmax
	lmax := make([]int, len(input))
	lmax[0] = input[0]
	for i := 1; i < len(input); i++ {
		lmax[i] = max(lmax[i-1], input[i])
	}

	// Calculate rmax
	rmax := make([]int, len(input))
	rmax[len(input)-1] = input[len(input)-1]
	for i := len(input) - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], input[i])
	}

	fmt.Println(lmax)
	fmt.Println(rmax)

	amount := 0
	for i := 1; i < len(input)-1; i++ {
		amount = amount + min(lmax[i], rmax[i]) - input[i]
	}

	return amount
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(TrapRainWater(a))
	fmt.Println(TrapRainWaterEff(a))

	a = []int{5, 4, 3}
	fmt.Println(TrapRainWater(a))
	fmt.Println(TrapRainWaterEff(a))

	a = []int{2, 0, 2}
	fmt.Println(TrapRainWater(a))
	fmt.Println(TrapRainWaterEff(a))
}
