package main

import "fmt"

func SelectionSort(input []int) []int {
	currentIndex := 0

	for currentIndex < len(input) {
		smallestIndex := currentIndex
		for i := currentIndex + 1; i < len(input); i++ {
			if input[smallestIndex] > input[i] {
				smallestIndex = i
			}
		}

		input[currentIndex], input[smallestIndex] = input[currentIndex], input[smallestIndex]
		currentIndex = currentIndex + 1
	}

	return input
}

type MaxHeap struct {
	Elements []int

	/*
		Add below attributes for static arrays

		size     int
		capacity int
	*/
}

func New() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(key int) {
	h.Elements = append(h.Elements, key)

	// Apply heapify operation on the last element
	h.MaxHeapifyUp(len(h.Elements) - 1)
}

func (h *MaxHeap) MaxHeapifyUp(index int) {
	for h.Elements[h.Parent(index)] < h.Elements[index] {
		h.Swap(h.Parent(index), index)
		index = h.Parent(index)
	}
}

func (h *MaxHeap) Swap(i1, i2 int) {
	h.Elements[i1], h.Elements[i2] = h.Elements[i2], h.Elements[i1]
}

func (h *MaxHeap) MaxHeapify(index int) { // Down version
	left := h.Left(index)
	right := h.Right(index)
	size := len(h.Elements)

	largest := index
	if left < size && h.Elements[left] > h.Elements[largest] {
		largest = left
	}

	if right < size && h.Elements[right] > h.Elements[largest] {
		largest = right
	}

	if largest != index {
		h.Swap(largest, index)
		h.MaxHeapify(largest)
	}
}

func (h *MaxHeap) MaxHeapifyWithSize(index int, resize int) { // Down version
	left := h.Left(index)
	right := h.Right(index)
	size := resize
	if size <= 0 {
		return
	}
	largest := index
	if left < size && h.Elements[left] > h.Elements[largest] {
		largest = left
	}

	if right < size && h.Elements[right] > h.Elements[largest] {
		largest = right
	}

	if largest != index {
		h.Swap(largest, index)
		h.MaxHeapifyWithSize(largest, size)
	}
}

func BuildHeapFromArray(input []int) *MaxHeap {
	h := New()
	for i := 0; i < len(input); i++ {
		h.Elements = append(h.Elements, input[i])
	}

	// Fetch the index of bottom - right most index
	for i := (len(input) - 2) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}

	return h
}

func (h *MaxHeap) Parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) Left(index int) int {
	return (2 * index) + 1
}

func (h *MaxHeap) Right(index int) int {
	return (2 * index) + 2
}

func HeapSort(input []int) *MaxHeap {
	size := len(input)
	h := BuildHeapFromArray(input)
	for i := size - 1; i >= 0; i-- {
		h.Swap(0, i)
		h.MaxHeapifyWithSize(0, i) // Reduce the size at each iteration
	}

	return h
}

func main() {
	input := []int{50, 20, 4, 10, 15}
	fmt.Println(HeapSort(input))
}
