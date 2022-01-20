package main

import (
	"container/heap"
)

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int    { return len(h) }
func (h MinHeap) Empty() bool { return len(h) == 0 }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() int           { return h[0] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxHeap) Len() int    { return len(h) }
func (h MaxHeap) Empty() bool { return len(h) == 0 }

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() int           { return h[0] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	elements []int
	minHeap  *MinHeap
	maxHeap  *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		elements: []int{},
		minHeap:  &MinHeap{},
		maxHeap:  &MaxHeap{},
	}
}

func (this *MedianFinder) InitializeHeaps() {
	heap.Init(this.minHeap)
	heap.Init(this.maxHeap)
}

func (this *MedianFinder) AddNum(num int) { // log(N)
	this.elements = append(this.elements, num)

	if this.maxHeap.Len() <= 0 || this.maxHeap.Top() > num {
		this.maxHeap.Push(num)
	} else {
		this.minHeap.Push(num)
	}

	if this.maxHeap.Len() > this.minHeap.Len()+1 {
		this.minHeap.Push(this.maxHeap.Top())
		this.maxHeap.Pop()
	} else if this.minHeap.Len() > this.maxHeap.Len()+1 {
		this.maxHeap.Push(this.minHeap.Top())
		this.minHeap.Pop()
	}
	// sort.Ints(this.elements)
}

func (this *MedianFinder) FindMedian() float64 {
	size := len(this.elements)

	// maxHeapHalf := this.elements[0:(size / 2)]
	// minHeapHalf := this.elements[(size)/2:]

	// for i := 0; i < len(maxHeapHalf); i++ {
	// 	heap.Push(this.maxHeap, maxHeapHalf[i])
	// }

	// for i := 0; i < len(minHeapHalf); i++ {
	// 	heap.Push(this.minHeap, minHeapHalf[i])
	// }

	var median float64

	if size%2 == 0 {
		median = float64(this.maxHeap.Top()+this.minHeap.Top()) * 0.5
	} else if size%2 == 1 {
		median = float64(this.maxHeap.Top())
	}

	return median
}

func main() {

}
