package main

/*
	Approach 1 :
		- Sort an array
		- Pick k largest elements from the end
		- O(NLOGN)
	Approach 2 :
		- Prepare a Binary Heap (Array can be represented as Binary heap - nothing to be done here)
		- Build a Max Heap
		- Extract Max until k
		- O(N + NLOGK)
	Approach 3:
		- Prepare Binary tree
		- Build a min heap for first k elements
		- Loop from k + 1 elements to the end of an array
			- While looping compare top of minheap with array value
			- If top of minheap is smaller we extract it and push the array value
			- Repeat until you reach the end of an array
		- In the End min heap will have 3 largest elements of a given input
		- O(K + (N-K)LOGK)
*/
func main() {

}
