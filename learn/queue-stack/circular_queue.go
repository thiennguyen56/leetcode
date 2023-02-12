package main

import "fmt"

type MyCircularQueue struct {
	data     []int
	capacity int
	head     int
	tail     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:     make([]int, k),
		capacity: k,
		head:     -1,
		tail:     -1,
	}
}

func (this *MyCircularQueue) EnQueue(v int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
	}
	this.tail = (this.tail + 1) % this.capacity
	this.data[this.tail] = v
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
	this.head = (this.head + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == -1
}

func (this *MyCircularQueue) IsFull() bool {
	return this.head == (this.tail+1)%this.capacity
}

func main() {
	fmt.Println("vim-go")
}
