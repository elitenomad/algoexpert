package main

import (
	"strconv"
	"strings"
)

func addStrings(num1 string, num2 string) string {
	p1 := len(num1) - 1
	p2 := len(num2) - 1

	carry := 0
	res := []int{}

	for p1 >= 0 || p2 >= 0 {
		s1 := 0
		s2 := 0
		if p1 >= 0 {
			s1 = int(num1[p1] - '0')
		}

		if p2 >= 0 {
			s2 = int(num2[p2] - '0')
		}

		sum := (s1 + s2 + carry) % 10
		carry = (s1 + s2 + carry) / 10
		res = append(res, sum)

		p1 -= 1
		p2 -= 1
	}

	if carry != 0 {
		res = append(res, carry)
	}

	res = reverseArray(res)

	ids := []string{}
	for _, i := range res {
		ids = append(ids, strconv.Itoa(i))
	}

	return strings.Join(ids, "")
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
