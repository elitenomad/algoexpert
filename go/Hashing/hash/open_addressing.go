package main

import "fmt"

// Hash represent the structure we will push elements to
// -1 key represents no key
// -2 represents its marked delete
type Hash struct {
	buckets  []int
	capacity int
}

func New(capacity int) *Hash {
	initialArray := make([]int, capacity)

	for index, _ := range initialArray {
		initialArray[index] = -1
	}

	return &Hash{
		buckets:  initialArray,
		capacity: capacity,
	}
}

func (h *Hash) hashKey(key int) int {
	return key % h.capacity
}

func (h *Hash) Add(key int) bool {
	// Fetch the key
	idx := h.hashKey(key)

	for h.buckets[idx] != -1 && h.buckets[idx] != -2 && h.buckets[idx] != key {
		idx = (idx + 1) % h.capacity
	}

	if h.buckets[idx] == key {
		return false
	} else {
		h.buckets[idx] = key
		return true
	}
}

func (h *Hash) Search(key int) bool {
	idx := h.hashKey(key)
	hdx := idx

	for h.buckets[hdx] != -1 {

		if h.buckets[hdx] == key {
			return true
		}

		hdx = (hdx + 1) % h.capacity

		if hdx == idx {
			return false
		}
	}
	return false
}

func (h *Hash) Delete(key int) bool {
	idx := h.hashKey(key)
	hdx := idx
	for h.buckets[hdx] != -1 {
		if h.buckets[idx] == key {
			h.buckets[idx] = -2
			return true
		}

		hdx = (hdx + 1) % h.capacity
		if hdx == idx {
			return false
		}
	}

	return false
}

func main() {
	h := New(31)
	h.Add(10)
	fmt.Println("Hash value :", h)
	h.Add(20)
	fmt.Println("Hash value :", h)
	h.Add(30)
	fmt.Println("Hash value :", h)
	h.Search(20)
	fmt.Println("Hash value :", h)
	h.Search(10)
	fmt.Println("Hash value :", h)
	h.Delete(30)
	fmt.Println("Hash value :", h)
	h.Search(30)
	fmt.Println("Hash value :", h)
}
