package main

import "testing"

type MyStack struct {
	queue []int
}

func Constructor() *MyStack {
	return &MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.queue = this.queue[:len(this.queue)-1]
	return top
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func TestMyStack(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	t.Log(obj.Top())
	t.Log(obj.Empty())
	t.Log(obj.Pop())
	t.Log(obj.Empty())
}
