package main

// Do not edit the class below except for the
// PopulateSuffixTrieFrom and Contains methods.
// Feel free to add new properties and methods
// to the class.
type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {

	for j := range str {
		node := trie
		for i := j; i < len(str); i++ {
			char := str[i]
			if node[char] == nil {
				node[char] = NewSuffixTrie()
			}

			node = node[char]
		}
		node['*'] = nil
	}
}

func (trie SuffixTrie) Contains(str string) bool {
	node := trie

	for i := 0; i < len(str); i++ {
		char := str[i]
		if _, found := node[char]; !found {
			return false
		}

		node = node[char]
	}

	_, found := node['*']

	return found
}
