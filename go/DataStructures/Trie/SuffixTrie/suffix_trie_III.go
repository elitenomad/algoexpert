package main

import "fmt"

// NOT WORKING
type TrieNode struct {
	value     map[byte]*TrieNode
	isWordEnd bool
}

type Trie struct {
	root *TrieNode
}

func InitializeTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(word string) {
	for i := range word {
		currentNode := t.root
		for j := i; j < len(word); j++ {
			char := word[j]
			if currentNode.value[char] == nil {
				currentNode.value = make(map[byte]*TrieNode)
				currentNode.value[char] = &TrieNode{}
			}
			currentNode = currentNode.value[char]
		}
		currentNode.isWordEnd = true
	}
}

func (t *Trie) Contains(word string) bool {
	currentNode := t.root
	for j := 0; j < len(word); j++ {
		char := word[j]

		if currentNode.value[char] == nil {
			return false
		}

		currentNode = currentNode.value[char]
	}

	if currentNode.isWordEnd {
		return true
	}

	return false
}

func main() {
	trie := InitializeTrie()
	trie.Insert("babc")

	fmt.Println(trie.Contains("babc"))
	fmt.Println(trie.Contains("c"))
	fmt.Println(trie.Contains("bca"))
	fmt.Println(trie.Contains("abc"))
	fmt.Println(trie.Contains("zoo"))
}
