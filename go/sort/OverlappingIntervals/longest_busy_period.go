package main

import "fmt"

func LongestBusyPeriod(input []int) int {
	longestPeriod := 0
	slotSet := map[int]bool{}

	for i := 0; i < len(input); i++ {
		slotSet[input[i]] = true
	}

	for slot, _ := range slotSet {
		currentSlot := slot
		currentConsecutiveSlot := 1
		_, temp := slotSet[currentSlot+1]
		for temp {
			currentSlot++
			currentConsecutiveSlot++
			_, temp = slotSet[currentSlot+1]
		}

		longestPeriod = max(longestPeriod, currentConsecutiveSlot)
	}

	return longestPeriod
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	schedule := []int{3, 1, 15, 5, 2, 12, 10, 4, 8, 9}
	fmt.Println(LongestBusyPeriod(schedule))
	schedule = []int{3, 1, 12, 5, 2, 11, 10, 4, 8, 9, 13, 14, 15}
	fmt.Println(LongestBusyPeriod(schedule))
}
