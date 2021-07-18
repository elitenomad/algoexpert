package main

import "fmt"

func cut_the_rope(n, a, b, c int) int {
	return ctRoleHelper_Tutor_version(n, a, b, c)
}
func ctRoleHelper_Tutor_version(n, a, b, c int) int { // O(3^n)
	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	c1 := ctRoleHelper_Tutor_version(n-a, a, b, c)
	c2 := ctRoleHelper_Tutor_version(n-b, a, b, c)
	c3 := ctRoleHelper_Tutor_version(n-c, a, b, c)

	max := -1
	if c1 >= c2 && c1 >= c3 {
		max = c1
	}

	if c2 >= c3 && c2 >= c1 {
		max = c2
	}

	if c3 >= c1 && c3 >= c2 {
		max = c3
	}

	if max == -1 {
		return -1
	}

	return max + 1
}

func ctRoleHelper(n, a, b, c, d int) int {
	if n == 0 {
		return d
	}

	if n < 0 {
		return -1
	}

	c1 := ctRoleHelper(n-a, a, b, c, d+1)
	c2 := ctRoleHelper(n-b, a, b, c, d+1)
	c3 := ctRoleHelper(n-c, a, b, c, d+1)

	max := -1
	if c1 >= c2 && c1 >= c3 {
		max = c1
	}

	if c2 >= c3 && c2 >= c1 {
		max = c2
	}

	if c3 >= c1 && c3 >= c2 {
		max = c3
	}

	return max
}

func main() {
	fmt.Println(cut_the_rope(23, 11, 9, 12)) // Expected is 2 :)
	fmt.Println(cut_the_rope(5, 2, 5, 1))    // Expected is 5 :)
	fmt.Println(cut_the_rope(9, 2, 2, 2))    // Expected is 5 :)
}
