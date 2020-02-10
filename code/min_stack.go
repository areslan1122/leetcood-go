package main

import "fmt"

type MinStack struct {
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{
		minStack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.minStack = append(this.minStack, x)
}

func (this *MinStack) Pop() {
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {

	return this.minStack[len(this.minStack)-1]
}

func (this *MinStack) GetMin() int {

	min := this.minStack[0]
	for i := 0; i < len(this.minStack); i++ {

		if min > this.minStack[i] {
			min = this.minStack[i]
		}
	}

	return min

}

func main() {

	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_1 := obj.GetMin()
	obj.Pop()
	param_2 := obj.Top()
	param_3 := obj.GetMin()

	fmt.Println(param_1, param_2, param_3)

}
