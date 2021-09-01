package hash

import (
	"fmt"
	"hash/maphash"
)

const (
	numOfBuckets = 256
)

type entry struct {
	key   string
	value int
}

type Hash struct {
	buckets [][]entry
	hash    maphash.Hash
}

func New() *Hash {
	return &Hash{
		buckets: make([][]entry, numOfBuckets),
	}
}

func (h *Hash) hashKey(key string) int {
	// Reset mphash to fetch the latest value everytime you
	// send a key will receive same value
	h.hash.Reset()

	h.hash.WriteString(key)

	n := h.hash.Sum64()

	return int(n % numOfBuckets)
}

// Add or Store a value to hash based on a key
func (h *Hash) Store(key string, value int) {
	// Find the key where the key must go
	idx := h.hashKey(key)

	// Fetch the bucket we should place entry{k, v}
	bucket := h.buckets[idx]

	// Usecase 1: what if the value already exist
	for idx := range bucket {
		if bucket[idx].key == key {
			bucket[idx].value = value
			return
		}
	}

	// Usecase 2, its a new entry and we need to
	// Append the k, v to bucket idx array
	h.buckets[idx] = append(bucket, entry{key, value})
}

func (h *Hash) Retrive(key string) (int, error) {
	// Find the idx of by the key
	idx := h.hashKey(key)

	// Fetch the bucket
	bucket := h.buckets[idx]

	// Loop through bucket entries and find value
	for id := range bucket {
		if bucket[id].key == key {
			return bucket[id].value, nil
		}
	}

	// Value do not exist
	return -1, fmt.Errorf("%q not found", key)
}

func (h *Hash) Delete(key string) error {
	//Fetch the idx
	idx := h.hashKey(key)

	// Extract a copy of the bucket from the hash table.
	bucket := h.buckets[idx]

	// Iterate over the entries for the specified bucket.
	for entryIdx, entry := range bucket {

		// Compare the keys and if there is a match remove
		// the entry from the bucket.
		if entry.key == key {

			// Remove the entry based on its index position.
			bucket = // Need to delete entry from buckets removeEntry(bucket, entryIdx)

			// Replace the existing bucket for the new one.
			h.buckets[bucketIdx] = bucket
			return nil
		}
	}

	// The key was not found so return the error.
	return fmt.Errorf("%q not found", key)
}

func removeEntry(bucket []entry, idx int) []entry {

	// https://github.com/golang/go/wiki/SliceTricks
	copy(bucket[idx:], bucket[idx+1:])

	// Set the proper length for the new slice since
	// an entry was removed. The length needs to be
	// reduced by 1.
	bucket = bucket[:len(bucket)-1]

	// Look to see if the current allocation for the
	// bucket can be reduced due to the amount of
	// entries removed from this bucket.
	return reduceAllocation(bucket)
}

// reduceAllocation looks to see if memory can be freed to
// when a bucket has lost a percent of entries.
func reduceAllocation(bucket []entry) []entry {

	// If the bucket if more than ½ full, do nothing.
	if cap(bucket) < 2*len(bucket) {
		return bucket
	}

	// Free memory when the bucket shrinks a lot. If we don't do that,
	// the underlying bucket array will stay in memory and will be in
	// the biggest size the bucket ever was
	newBucket := make([]entry, len(bucket))
	copy(newBucket, bucket)
	return newBucket
}

func (h *Hash) Len() int {
	// Initialize a sum variable
	sum := 0

	// Loop through all the buckets
	for _, bucket := range h.buckets {
		// Calculate the length of secondary array
		sum += len(bucket)
	}

	// Return the total length
	return sum
}
