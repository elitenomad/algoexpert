package main

import "fmt"

const (
	numOfBuckets = 7 // Going with this number
)

// Hashmap structure
type HashMap struct {
	buckets [numOfBuckets]*Bucket
}

// Bucket structure
type Bucket struct {
	Head *BucketNode
}

type BucketNode struct {
	Key   string
	Value int
	Next  *BucketNode
}

func (h *HashMap) fetchHashKey(key string) int {
	sum := 0
	for i := 0; i < len(key); i++ {
		sum += int(key[i])
	}
	return sum % len(h.buckets)
}

// Hash table
// Insert
func (h *HashMap) Insert(key string) {
	// Find the hash key
	hashKey := h.fetchHashKey(key)
	// Figure out which bucket it will be part of
	bucket := h.buckets[hashKey]
	// Place it in head of the bucket
	bucket.Insert(key)
}

// Search

func (h *HashMap) Search(key string) bool {
	// Find the hash key
	hashKey := h.fetchHashKey(key)
	// Lookup the right bucket
	bucket := h.buckets[hashKey]

	// search the value the linked list
	return bucket.Search(key)
}

// Delete - returns the key
func (h *HashMap) Delete(key string) string {
	// Find the hash Key
	hashKey := h.fetchHashKey(key)

	// Lookup the right bucket
	bucket := h.buckets[hashKey]

	// Call the delete on the Linked List structure
	deletedKey := bucket.Delete(key)

	return deletedKey
}

// Buckets
// Insert
func (b *Bucket) Insert(key string) {
	bucket := &BucketNode{Key: key, Next: b.Head}
	b.Head = bucket
}

// Search
func (b *Bucket) Search(key string) bool {
	current := b.Head

	for current != nil {
		if current.Key == key {
			return true
		}
		current = current.Next
	}

	return false
}

// Delete
func (b *Bucket) Delete(key string) string {
	current := b.Head
	deletedKey := ""

	if current == nil {
		return deletedKey
	}

	if current.Next == nil && current.Key == key {
		b.Head = nil
		return deletedKey
	}

	for current != nil {
		if current.Next.Key == key {
			deletedKey = key
			current.Next = current.Next.Next
		}
	}

	return deletedKey
}

// Simple Hash function

// Init method to initialize DS
func InitHashMap() *HashMap {
	h := &HashMap{
		buckets: [numOfBuckets]*Bucket{},
	}

	for i := 0; i < len(h.buckets); i++ {
		h.buckets[i] = &Bucket{}
	}

	return h
}

func main() {
	h := InitHashMap()
	fmt.Println(h)
	fmt.Println()

	h.Insert("abc")
	h.Insert("xyz")
	h.Insert("pranava")
	h.Insert("swaroop")
	fmt.Println(h.Search("pranava"))
	h.Delete("xyz")
	fmt.Println(h.Search("xyz"))
	fmt.Println(h.Search("swaroop"))
}
