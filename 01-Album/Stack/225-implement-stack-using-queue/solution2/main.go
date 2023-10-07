package solution2


type MyStack struct {
	q1 []int
	q2 []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		q1: make([]int, 0),
		q2: make([]int, 0),

	}
}

// push to back
// peek/pop from front, 有这个特性就好了啊


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.q1 = append(this.q1, x)

}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	// swap q1 with q2 to avoid copying all elements from q2 to q1.

	for len(this.q1) > 1 {
		this.q2 = append(this.q2, this.q1[0])
		this.q1 = this.q1[1:]
	}
	ret := this.q1[0]
	this.q1 = this.q2
	this.q2 = []int{}
	return ret
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q1[len(this.q1)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
