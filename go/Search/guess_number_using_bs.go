/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
package main

func guessNumber(n int) int {
	l := 1
	r := n

	for l <= r {
		mid := l + (r-l)/2
		guessy := guess(mid) //This method is from API returns Leetcode problem

		if guessy == 0 {
			return mid
		}

		if guessy < 0 {
			r = mid - 1
		} else if guessy > 0 {
			l = mid + 1
		}
	}

	return -1
}
