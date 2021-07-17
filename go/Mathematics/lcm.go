package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/*
	LCM of two numbers is bigger than each of the numbers
	Lets start with a max of both numbers
	Initialize one number,
	Keep incrementing and find the multiples
*/
func Lcm(a, b int) int {
	mx := max(a, b)
	n := 1

	for true {
		if mx%a == 0 && mx%b == 0 {
			break
		}

		mx *= n
		n += 1
	}

	return mx
}

//Approach - 3 Optimized version of EA
func gcd_3(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd_3(b, a%b)
}

// Approach 2 would be mathematical
// a * b = GCD / LCM , LCM = GCD / a * b
func Lcm_2(a, b int) int {
	lcmey := (a * b) / gcd_3(a, b)
	return lcmey
}

func main() {
	a, b := 10, 15
	fmt.Println(Lcm(a, b))
	fmt.Println(Lcm_2(a, b))
}
