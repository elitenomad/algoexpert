package main

import "fmt"

type MyCircularQueue struct {
	Elements []interface{}
	size     int
	head     int
	tail     int
}

func Constructor(k int) MyCircularQueue {
	q := MyCircularQueue{}
	q.Elements = make([]interface{}, k)
	q.head = -1
	q.tail = -1
	q.size = k

	return q
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.head = 0
	}

	this.tail = (this.tail + 1) % this.size
	this.Elements[this.tail] = value

	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}

	this.head = (this.head + 1) % this.size
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Elements[this.head].(int)
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Elements[this.tail].(int)
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return ((this.tail + 1) % this.size) == this.head
}

// Your MyCircularQueue object will be instantiated and called as such:
func main() {
	k := 10
	obj := Constructor(k)
	obj.EnQueue(1)
	obj.EnQueue(3)
	obj.EnQueue(14)
	obj.EnQueue(2)
	obj.EnQueue(10)

	param_2 := obj.DeQueue()
	fmt.Println("What is param_2 ", param_2)

	param_3 := obj.Front()
	fmt.Println("What is param_3 ", param_3)

	param_4 := obj.Rear()
	fmt.Println("What is param_4 ", param_4)

	param_5 := obj.IsEmpty()
	fmt.Println("What is param_5 ", param_5)

	param_6 := obj.IsFull()
	fmt.Println("What is param_6 ", param_6)

}
