package main

import "fmt"

/*
	A trie is a data structure for efficient information retrieval.
	It is a special kind of tree where a path starting from root to a
	particular node can define a word that is stored in this tree.
	A trie can be built for entire ASCII_SIZE, ALPHABETS,  NUMBERS
	depending upon the use case. For, eg below will be the properties
	of a Trie data structure supporting small case alphabet letters

	Each node has ALBHABET_SIZE=26 children. Each child node is a trie
	node itself and has ALBHABET_SIZE=26 children.Each node is at an
	index in its parent’ children node array and represents an ASCII
	character. For eg for a particular node, the first non-nil children
	node will mean the presence of char ‘a’,  second non-nil children
	node means the presence of ‘b’ and so on. Absence of a child at
	an index means no value Each node also has a boolean field indicating
	whether the node is the end of word or not. The root node is the
	starting node and has ALBHABET_SIZE=26 children. root is associated with an empty value
*/

const (
	// TOTAL CHARS in English Alphabets
	ALPHABET_SIZE = 26
)

type trieNode struct {
	childrens [ALPHABET_SIZE]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) Insert(word string) {
	wordLength := len(word)
	current := t.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		// fmt.Println("INDEX is :", index)
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) PrintCount(root *trieNode, count int) {
	if root.isWordEnd {
		fmt.Println(count)
		return
	}
	for i := 0; i < ALPHABET_SIZE; i++ {
		if root.childrens[i] != nil {
			t.PrintCount(root.childrens[i], count+1)
		}
	}
}

func (t *trie) Search(word string) bool {
	wordLength := len(word)
	current := t.root

	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}

		current = current.childrens[index]
	}

	return current != nil && current.isWordEnd
}

func (t *trie) Delete(root *trieNode, word string, i int) *trieNode {
	if root == nil {
		return nil
	}

	if i == len(word) {
		root.isWordEnd = false
		if t.IsEmpty(root) {
			root = nil
			return root
		}
	}

	index := word[i] - 'a'
	root.childrens[index] = t.Delete(root.childrens[index], word, i+1)

	if t.IsEmpty(root) && root.isWordEnd == false {
		root = nil
	}

	return root
}

func (t *trie) IsEmpty(root *trieNode) bool {
	for i := 0; i < ALPHABET_SIZE; i++ {
		if root.childrens[i] != nil {
			return false
		}
	}
	return true
}

func main() {
	trie := initTrie()
	output := []string{"Yes", "No"}
	words := []string{"sam", "john", "tim", "jose", "rose", "cat", "dog", "dogg", "roses"}
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}

	// trie.PrintCount(trie.root, 0)

	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.Search(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}

	trie.Insert("zoo")
	fmt.Println("FOUND ZOO: ", trie.Search("zoo") == true)
	trie.Delete(trie.root, "zoo", 0)
	if trie.Search("zoo") == true {
		fmt.Println("Found it: ", output[0])
	} else {
		fmt.Println("Found it: ", output[1])
	}
}
