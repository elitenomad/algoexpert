package main

import (
	"fmt"
)

// Capacity dictates size of cache
// Size represents current size of cache
// minFreq placeholder to remove latest minFreq
// FreqDist =>
// keyDict =>
type LFUCache struct {
	capacity int
	size     int
	minFreq  int

	freqDict map[int]*LinkedList // 1 => 1 2 3 4
	keyDict  map[int]*LinkedListNode
}

func newLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		freqDict: make(map[int]*LinkedList),
		keyDict:  make(map[int]*LinkedListNode),
	}
}

func (lf *LFUCache) Get(key int) *LinkedListNode {
	if _, found := lf.keyDict[key]; !found {
		return nil
	}

	temp := lf.keyDict[key]

	// Fetch the Key from KeyDict
	// Delete Node from LinkedList for the freq of LinkedListNode
	// If that freq deletion in the above step makes linked list nil, remove freq key
	// if keyDict freq and minFreq are same increase minFreq
	lf.freqDict[temp.freq].DeleteNode(temp)
	if lf.freqDict[lf.keyDict[key].freq].head == nil {
		delete(lf.freqDict, lf.keyDict[key].freq)
		if lf.minFreq == lf.keyDict[key].freq {
			lf.minFreq++
		}
	}

	// increase the frequence of the key
	lf.keyDict[key].freq++
	if _, ok := lf.freqDict[lf.keyDict[key].freq]; !ok {
		lf.freqDict[lf.keyDict[key].freq] = &LinkedList{}
	}
	lf.freqDict[lf.keyDict[key].freq].Append(lf.keyDict[key])
	return lf.keyDict[key]
}

func (lf *LFUCache) Set(key, value int) {
	// Use case when the value exist and need to update the value
	if lf.Get(key) != nil {
		lf.keyDict[key].val = value
		return
	}

	// Usecase when Size == Capacity
	if lf.size == lf.capacity {
		// Delete least freq used one from keyDict and freqDict
		minFreqNode := lf.freqDict[lf.minFreq].head
		delete(lf.keyDict, minFreqNode.key)
		lf.freqDict[lf.minFreq].DeleteNode(minFreqNode)

		//
		if lf.freqDict[lf.minFreq].head == nil {
			delete(lf.freqDict, lf.minFreq)
		}

		lf.size -= 1
	}

	// set the minFreq = 1
	lf.minFreq = 1
	lf.keyDict[key] = &LinkedListNode{key: key, val: value, freq: lf.minFreq}

	// Add a linkedList if the freq key that do not exist
	if _, found := lf.freqDict[lf.keyDict[key].freq]; !found {
		lf.freqDict[lf.keyDict[key].freq] = &LinkedList{}
	}

	// Recently added node is appendd to freqDict linkedList
	lf.freqDict[lf.keyDict[key].freq].Append(lf.keyDict[key])
	lf.size += 1
}

func (lf *LFUCache) PrintDict() {
	for first, second := range lf.keyDict {
		fmt.Printf("(%v,%v)", first, second.val)
	}
	fmt.Println("")
}

func (lf *LFUCache) Print() string {
	str := ""
	for key, data := range lf.keyDict {
		str += fmt.Sprintf("(%v,%v)", key, data.val)
	}
	return str
}
