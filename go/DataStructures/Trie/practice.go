package main

import "fmt"

// Because we are dealing with array of words
// Words are made of alphabets . we have only 26
// Trie structure we deal wth will nodes which has
// 26 children and so on
const (
	ALBHABET_SIZE = 26
)

// Create a TrieNode structure
type TrieNode struct {
	children  [ALBHABET_SIZE]*TrieNode
	isWordEnd bool
}

// Create a Trie structure
type Trie struct {
	root *TrieNode
}

// Initialize the Trie structure to create dictionary of words
func InitializeTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

// Operations we want to implement on that are
// Insert
// Search
// Delete
func (t *Trie) Insert(word string) {
	currentNode := t.root

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a' // Technically gives you the array indexes from 0-26
		if currentNode.children[index] == nil {
			currentNode.children[index] = &TrieNode{}
		}
		currentNode = currentNode.children[index]
	}

	currentNode.isWordEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a' // Provides the array index starting 0-25

		if currentNode.children[index] == nil {
			return false
		}

		currentNode = currentNode.children[index]
	}

	return currentNode != nil && currentNode.isWordEnd
}

func (t *Trie) Delete(root *TrieNode, word string, i int) *TrieNode {
	if root == nil {
		return nil
	}

	if i == len(word) {
		root.isWordEnd = false
		if t.IsEmpty(root) {
			return nil
		}
	}

	index := word[i] - 'a'
	root.children[index] = t.Delete(root.children[index], word, i+1)

	if t.IsEmpty(root) && root.isWordEnd == false {
		root = nil
	}

	return root
}

func (t *Trie) IsEmpty(root *TrieNode) bool {
	for i := 0; i < ALBHABET_SIZE; i++ {
		if root.children[i] != nil {
			return false
		}
	}

	return true
}

func main() {
	trie := InitializeTrie()
	fmt.Println(trie.root)
	trie.Insert("kurnool")
	fmt.Println(trie.root.children[10])
	trie.Insert("geeks")
	trie.Insert("geek")
	fmt.Println(trie.Search("geek"))
	fmt.Println(trie.Search("gee"))
}
