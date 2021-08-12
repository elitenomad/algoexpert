func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) < len(sequence) {
		return false
	}

	seq := 0
	i_len := 0

	for seq < len(sequence) && i_len < len(array) {
		if sequence[seq] == array[i_len] {
			seq = seq + 1
		}

		i_len = i_len + 1
	}

	return seq == len(sequence)
}

func main() {
	a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, 10}
	fmt.Println(IsValidSubsequence(a, seq))
}