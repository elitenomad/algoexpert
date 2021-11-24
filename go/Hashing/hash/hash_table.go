package main

import "fmt"

const (
	NUM_SIZE = 7
)

// Hash table structure
type Hash struct {
	Buckets [NUM_SIZE]*Bucket
}

// Bucket structure
type Node struct {
	Key  string
	Next *Node
}

type Bucket struct {
	Head *Node
}

// Insert
func (h *Hash) Insert(key string) {
	// Find hash
	k := Hashing(key)

	// Find bucket
	bucket := h.Buckets[k]

	// Insert into bucket
	bucket.Insert(key)
}

// Delete
func (h *Hash) Delete(key string) {
	// Find hash
	k := Hashing(key)

	// Find bucket
	bucket := h.Buckets[k]

	// Insert into bucket
	bucket.Delete(key)
}

// Search
func (h *Hash) Search(key string) bool {
	// Find hash
	k := Hashing(key)

	// Find bucket
	bucket := h.Buckets[k]

	// Insert into bucket
	return bucket.Search(key)
}

// Insert
func (b *Bucket) Insert(key string) {
	if b.Search(key) {
		return
	}
	node := &Node{Key: key, Next: b.Head}
	b.Head = node
}

// Search
func (b *Bucket) Search(key string) bool {
	curr := b.Head

	for curr != nil {
		if curr.Key == key {
			return true
		}
		curr = curr.Next
	}

	return false
}

// Delete
func (b *Bucket) Delete(key string) {
	curr := b.Head

	// Consider use cases when key is head
	if curr.Key == key {
		b.Head = b.Head.Next
		return
	}
	// Consider use case when key is tail

	for curr != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			break
		}

		curr = curr.Next
	}
}

// hash
func Hashing(key string) int {
	sum := 0

	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}

	return sum % NUM_SIZE
}

// Init
func Init() *Hash {
	h := &Hash{}

	for i := 0; i < len(h.Buckets); i++ {
		h.Buckets[i] = &Bucket{}
	}

	return h
}

func main() {
	h := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		h.Insert(v)
	}
	fmt.Println(h.Search("STAN"))

	h.Delete("STAN")
	fmt.Println(h.Search("STAN"))
}
