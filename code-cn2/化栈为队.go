package main

type MyQueue struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		arr: []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	res := this.arr[0]
	this.arr = this.arr[1:]

	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.arr[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.arr) == 0 {
		return true
	} else {
		return false
	}

}
