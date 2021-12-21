package main

import "fmt"

/*
	Max Heap:
		Complete Binary Tree
		Can be represented in the form of an array
		left, right and parent : 2i-1, 2i+2 and (i-1)/2-floor
		Node value must be greater than its descenadants

	Operations:
		Insert
		Delete
		Heapify
*/

// Max heap struct will have a slice that holds the elements
type MaxHeap struct {
	Elements []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.Elements = append(h.Elements, key)
	h.MaxHeapifyUp(len(h.Elements) - 1)
}

// MaxHeapifyUp Operation
func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.Elements[parent(index)] < h.Elements[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Extract returns the largest key and removed it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.Elements[0]
	l := len(h.Elements) - 1
	if l <= 0 {
		return -1
	}
	h.Elements[0] = h.Elements[l]
	h.Elements = h.Elements[:l]

	h.MaxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.Elements) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		// When left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.Elements[l] > h.Elements[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.Elements[index] < h.Elements[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return (2 * i) + 1
}

func right(i int) int {
	return (2 * i) + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.Elements[i1], h.Elements[i2] = h.Elements[i2], h.Elements[i1]
}

func main() {
	h := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		h.Insert(v)
		fmt.Println(h.Elements)
	}

	for i := 0; i < 5; i++ {
		h.Extract()
		fmt.Println(h.Elements)
	}
}
