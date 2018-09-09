package main

import "fmt"

type MinStack struct {
	minVal *int
	stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nil, make([]int , 0)}
}


func (this *MinStack) Push(x int)  {
	if this.minVal == nil || *this.minVal > x {
		this.minVal = &x
	}

	this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack) - 1]
		if len(this.stack) > 0 {
			*this.minVal = this.stack[0]
			for _, v := range this.stack {
				if *this.minVal > v {
					*this.minVal = v
				}
			}
		} else {
			this.minVal = nil
		}
	}
}


func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack) - 1]
	}

	return 0
}


func (this *MinStack) GetMin() int {
	return *this.minVal
}

/*
["MinStack","push","push","push","push","getMin","pop","getMin","pop","getMin","pop","getMin"]
[[],[2],[0],[3],[0],[],[],[],[],[],[],[]]
 */
func main() {
	stack := Constructor()
	stack.Push(2)
	stack.Push(0)
	stack.Push(3)
	stack.Push(0)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
}
