package main

/*
	suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

	[4,5,6,7,0,1,2] if it was rotated 4 times.
	[1,2,4,5,6,7,0] if it was rotated 7 times.

	Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results
	in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

	Given the sorted rotated array nums of unique elements,
	return the minimum element of this array.

	You must write an algorithm that runs in O(log n) time.
*/

func MinInSortedArray(input []int) int {
	l := 0
	h := len(input) - 1

	// Use case 1
	if len(input) == 1 {
		return input[0]
	}

	// Use case 2
	if len(input) == 0 {
		return -1
	}

	// Use case 3
	if input[l] < input[h] {
		return input[l]
	}

	for l < h {
		m := l + (h-l)/2

		if input[m-1] > input[m] && input[m] < input[m+1] {
			return input[m]
		}

		if input[m] < input[l] {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}

func main() {
}
