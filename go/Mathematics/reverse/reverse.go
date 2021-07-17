package reverse

func Reverse(n int) int {
	last := 0

	for n != 0 {
		r := n % 10

		last = last * 10

		last = last + r

		n = n / 10
	}

	return last
}
