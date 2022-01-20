package main

import "fmt"

type TrieNode struct {
	children  [26]*TrieNode
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
			index := word[j] - 'a'
			if currentNode.children[index] == nil {
				currentNode.children[index] = &TrieNode{}
			}
			currentNode = currentNode.children[index]
		}
		currentNode.isWordEnd = true
	}
}

func (t *Trie) Contains(word string) bool {
	currentNode := t.root
	for j := 0; j < len(word); j++ {
		index := word[j] - 'a'

		if currentNode.children[index] == nil {
			return false
		}
		currentNode = currentNode.children[index]
	}
	if currentNode.isWordEnd {
		return true
	}

	return false
}

func main() {
	trie := InitializeTrie()
	trie.Insert("babc")
	fmt.Println(trie.root)
	fmt.Println(trie.Contains("c"))
	fmt.Println(trie.Contains("bca"))
	fmt.Println(trie.Contains("abc"))
	fmt.Println(trie.Contains("zoo"))
}
