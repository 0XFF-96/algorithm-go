package main

// 155. Min Stack
type MinStack struct {
	arr []int
	min []int
}

func Constructor() MinStack {
	return MinStack{
		arr:[]int{},
		min:[]int{},
	}
}

func (this *MinStack) Push(x int)  {
	if len(this.min) == 0 || x <= this.GetMin() {
		this.min = append(this.min, x)
	}
	this.arr = append(this.arr, x)
}

func (this *MinStack) Pop()  {
	if this.arr[len(this.arr)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
