package main

import "fmt"

func Union(a []int, b []int) []int {
	result := []int{}

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}

		if j > 0 && b[j] == b[j-1] {
			j++
			continue
		}
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else if a[i] > b[j] {
			result = append(result, b[j])
			j++
		} else {
			result = append(result, a[i])
			i++
			j++
		}
	}

	for i < len(a) {
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}

		result = append(result, a[i])
		i++
	}

	for j < len(b) {
		if j > 0 && b[j] == b[j-1] {
			j++
			continue
		}
		result = append(result, b[j])
		j++
	}

	return result
}

func main() {
	a := []int{3, 5, 8, 30, 30, 30, 40, 40, 50, 50, 60, 60} //{2, 20, 20, 40, 60}
	b := []int{2, 2, 2, 2, 2, 8, 9, 10, 15}                 //{10, 20, 20, 20}
	fmt.Println(Union(a, b))
}
