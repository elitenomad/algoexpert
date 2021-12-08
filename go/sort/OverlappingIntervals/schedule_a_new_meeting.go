package main

import (
	"fmt"
	"sort"
)

func ScheduleAMeetingIntoAlreadyMeetings(input [][]int, meeting []int) [][]int {
	output := [][]int{}

	// Sort the timings
	sort.Slice(input, func(i, j int) bool {
		return input[i][0] < input[j][0]
	})

	i := 0
	n := len(input)

	// Append all the meetings that start before new meeting
	for i < n && meeting[0] > input[i][0] {
		output = append(output, input[i])
		i += 1
	}

	size := len(output)
	if size == 0 || output[size-1][1] < meeting[0] {
		output = append(output, meeting)
	} else {
		output[size-1][1] = max(output[size-1][1], meeting[1])
	}

	for i < n {
		size = len(output)
		if output[size-1][1] < input[i][0] {
			output = append(output, input[i])
		} else {
			output[size-1][1] = max(output[size-1][1], input[i][1])
		}
		i += 1
	}

	return output
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	meetingTimes := [][]int{{1, 3}, {4, 6}, {8, 10}, {10, 12}, {13, 15}, {16, 18}}
	newMeeting := []int{9, 13}

	fmt.Println(ScheduleAMeetingIntoAlreadyMeetings(meetingTimes, newMeeting))

	meetingTimes = [][]int{{1, 3}, {4, 6}, {8, 10}, {10, 12}, {13, 15}, {16, 18}}
	newMeeting2 := []int{19, 20}
	fmt.Println(ScheduleAMeetingIntoAlreadyMeetings(meetingTimes, newMeeting2))

}
