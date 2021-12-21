package main

import (
	"fmt"
	"math"
)

/*
	Min Heap:
		Lowest value item remains at the top - Highest priority is assigned to lowest value
		Complete Binary Tree
		Can be represented in the form of an array (having left right parent relationsships)
		left, right and parent : 2i+1, 2i+2 and (i-1)/2-floor
			- continuguos locations
			- cache friendly
		Node value must be less than its descenadants

	Operations:
		Insert
		Delete
		Heapify
*/

// Max heap struct will have a slice that holds the elements
type MinHeap struct {
	Elements []int

	// We can go with size and capacity
	// For now using the dynamic slice
	// size int
	// capacity int
}

type PriorityQueue struct {
	h *MinHeap
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		h: &MinHeap{},
	}
}

func (p *PriorityQueue) Add(value int) {
	p.h.Insert(value)
}

func (p *PriorityQueue) Peek() int {
	return p.h.Elements[0]
}

func (p *PriorityQueue) Poll() int {
	removedElement := p.h.Elements[0]
	p.h.Delete(0)
	return removedElement
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.h.Elements) <= 0
}

func KSortedArray(input []int, k int) []int {
	// Implementation of Priority Queue is something to look for
	// Lets assume we have PQ implementation set
	pq := NewPriorityQueue() // Integers

	for i := 0; i <= k; i++ {
		pq.Add(input[i])
	}
	fmt.Println(pq.h.Elements)

	index := 0
	for i := k + 1; i < len(input); i++ {
		input[index] = pq.Poll()
		index += 1
		pq.Add(input[i])
	}

	for !pq.IsEmpty() {
		input[index] = pq.Poll()
		index += 1
	}

	return input
}

func New() *MinHeap {
	return &MinHeap{}
}

// Insert adds an element to the heap
// Algorithm
// 	- Add an element to the end of heap
//	- Perform MinHeapifyUp operation to move the element to the right place.
func (h *MinHeap) Insert(key int) {
	h.Elements = append(h.Elements, key)

	// Apply heapify operation on the last element
	h.MinHeapifyUp(len(h.Elements) - 1)
}

// Algorithm
//	- Change the value of the index
func (h *MinHeap) DecreaseKey(index, value int) {
	h.Elements[index] = value
	h.MinHeapifyUp(index)
}

func (h *MinHeap) Delete(index int) {
	// Call DecreaseKey(index, math.MinInt32) => Reaches the root
	h.DecreaseKey(index, math.MinInt32)
	// Then call ExtractMin on the root
	h.Extract()
}

func BuildHeapFromArray(input []int) *MinHeap {
	h := New()
	for i := 0; i < len(input); i++ {
		h.Elements = append(h.Elements, input[i])
	}

	// Fetch the index of bottom - right most index
	for i := (len(input) - 2) / 2; i >= 0; i-- {
		h.MinHeapify(i)
	}

	return h
}

// MinHeapifyUp Operation
// Algorithm
// 	- Start with an index
//	- Verify its parent is greater than descendant
//	- If so swap
// 	- Continue until the element reaches the spot
//		- Either element reaches the top and exits the loop
//		- Or element at current Index will be greater than parent (satisfies min heap rule)
func (h *MinHeap) MinHeapifyUp(index int) {
	for h.Elements[parent(index)] > h.Elements[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Deleting a value from a minimum heap requires two steps
// Decrease a key to math.MinInt32 (this reaches the root)
// Run ExtractMin

// Extract returns the smallest key and removed it from the heap
func (h *MinHeap) Extract() int {
	l := len(h.Elements) - 1
	if l <= 0 {
		return -1 // Return math.MinInt32
	}

	/*
		If there is a size and capacity declared
		we need to handle a condition where
		size = 1 and when ever we want to remove
		we do size-=1 which means we have to remove the top of the
		element

		if size == 1 {
			size-=1
			return h.Elements[0]
		}
	*/

	extracted := h.Elements[0]
	h.Elements[0] = h.Elements[l]
	h.Elements = h.Elements[:l]

	h.MinHeapify(0)
	return extracted
}

// Algorithm
// 	- Collect left and right indexes of given index
//	- if both values are greater or equal to root then we return (root is at its right place)
//	- compare left and right to the given index
//	- which is smaller will be smaller in the execution of method
//	- Swap the values of smaller and index given
//	- Recursively call the MinHeapify for smaller index until min heap is at its right place.
func (h *MinHeap) MinHeapify(index int) { // Generally this is down
	lt := left(index)
	rt := right(index)

	// if lt >= len(h.Elements) || rt >= len(h.Elements) {
	// 	return
	// }

	smallest := index
	if lt < len(h.Elements) && h.Elements[lt] < h.Elements[index] {
		smallest = lt
	}

	if rt < len(h.Elements) && h.Elements[rt] < h.Elements[smallest] {
		smallest = rt
	}

	if smallest != index {
		h.swap(smallest, index)
		h.MinHeapify(smallest)
	}
}

func (h *MinHeap) GetMin() int {
	return h.Elements[0]
}

// Non-recursive one
func (h *MinHeap) MinHeapifyDown(index int) {
	lastIndex := len(h.Elements) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.Elements[l] > h.Elements[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.Elements[index] > h.Elements[childToCompare] {
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

func (h *MinHeap) swap(i1, i2 int) {
	h.Elements[i1], h.Elements[i2] = h.Elements[i2], h.Elements[i1]
}

func main() {
	// h := &MinHeap{}

	// buildHeap := []int{20, 80, 200, 90, 100, 250, 500, 120}
	// for _, v := range buildHeap {
	// 	h.Insert(v)
	// 	fmt.Println(h.Elements)
	// }
	// fmt.Println("***")
	// for i := 0; i < 2; i++ {
	// 	// h.Extract()
	// 	fmt.Println(h.Elements)
	// }

	// fmt.Println("***")
	// h.DecreaseKey(3, 10)
	// fmt.Println(h.Elements)

	// input := []int{10, 5, 20, 2, 4, 8}
	// fmt.Println(BuildHeapFromArray(input))

	input := []int{15, 20, 4, 15, 50}
	fmt.Println(KSortedArray(input, 2))
}
