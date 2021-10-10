package main

import "fmt"

// Create a Hash table of counts for both
// - String and pattern
// Pre compute hash counts at each index
// Loop through the string and pattern match only
// If the hash counts match
func RabinKarpAlgorithm(str, pattern string) {
	// create a sample hash of relavent string chars
	h := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
	}

	// Find pattern hash count - precompute
	patternHashCount := 0
	for i := 0; i < len(pattern); i++ {
		patternHashCount += h[rune(pattern[i])]
	}

	// Precompute hashes at each index of the string
	precompute := map[int]int{}
	precompute[0] = h[rune(str[0])] + h[rune(str[1])] + h[rune(str[2])]
	for i := 1; i < len(str)-len(pattern)+1; i++ {
		precompute[i] = precompute[i-1] - h[rune(str[i-1])] + h[rune(str[i+len(pattern)-1])]
	}

	for i := 0; i < len(str)-len(pattern)+1; i++ {
		if precompute[i] == patternHashCount {
			j := 0
			for j = 0; j < len(pattern); j++ {
				if pattern[j] != str[i+j] {
					break
				}
			}

			if j == len(pattern) {
				fmt.Printf("%d \t", i)
			}
		}
	}
}

func main() {
	RabinKarpAlgorithm("abdabcbabc", "abc")
}
