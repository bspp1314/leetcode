package main

import (
	"fmt"
	"math"
)





/** initialize your data structure here. */
type MinStack struct {
	stack []int
	minValue int
	minValueNum int
}
func Constructor() MinStack {
	return MinStack{
			stack: make([]int,0,4),
			minValue: math.MaxInt32,
		}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack,x)
	if x < this.minValue {
		this.minValue = x
		this.minValueNum = 1
	}else if x == this.minValue {
		this.minValueNum++
	}
}


func (this *MinStack) Pop()  {
	if len(this.stack) == 0 {
		return
	}

	v := this.Top()
	this.stack = this.stack[:len(this.stack)-1]
	if v == this.GetMin() {
		this.minValueNum--
	}

	if len(this.stack) == 0  {
		this.minValueNum = math.MaxInt32
		this.minValue = 0
	}

	if this.minValueNum == 0 {
		this.minValue = this.stack[0]
		this.minValueNum = 1

		for i := 1; i < len(this.stack); i++ {
			if this.minValue == this.stack[i] {
				this.minValueNum++
			}else if this.minValue > this.stack[i] {
				this.minValue = this.stack[i]
				this.minValueNum = 1
			}
		}
	}
}


func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}

	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minValue
}

func main()  {
	//["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	fmt.Println(minStack.Top())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())

}

