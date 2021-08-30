package main

// func Anagram(s, a string) bool {

// 	return s == a
// }

// func getCharCount(str string) map[int]int {
// 	const charCount = make(map[int]int, 0)

// 	for i := 0; i < len(str); i++ {
// 		const rune = str[i]

// 		if charCount[rune] {
// 			charCount[rune] = 1
// 		} else {
// 			charCount[rune]++
// 		}
// 	}

// 	return charCount
// }

// func Anagram(s, a string) bool {
// 	if len(s) != len(a) {
// 		return false
// 	}

// 	const firstCharCount = getCharCount(s)
// 	const secondCharCount = getCharCount(a)

// 	for k, v := range secondCharCount {
// 		if firstCharCount[k] != secondCharCount[k] {
// 			return false
// 		}
// 	}

// 	return true
// }

func Anagram(s, a string) bool {
	/*
		A , B
		maskA and maskB

		loop through size of s
		# if s != a size return false
		#
		for i = 0 i < s i++ {
			pos = s[i] - 'a'
			maskA += (1 << pos)
		}

		for i = 0 ; i < s ; i++ {
			pos = a[i] - 'a'
			maskB += (1 << ps)
		}

		maskA == maskB
	*/

	return false
}

func main() {

}
