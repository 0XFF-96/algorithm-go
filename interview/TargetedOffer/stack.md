
- 最小栈
```
type MinStack struct {
    stack1 []int
    stack2 []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stack1: make([]int, 0),
        stack2: make([]int, 0),
    }
}


func (this *MinStack) Push(x int)  {
    this.stack1 = append(this.stack1, x)

    // 这里错了
    if len(this.stack2) == 0 ||  this.stack2[len(this.stack2)-1] >= x {
        this.stack2 = append(this.stack2, x)
    }
}


func (this *MinStack) Pop()  {
    val := this.stack1[len(this.stack1)-1]
    this.stack1 = this.stack1[:len(this.stack1)-1]
    if val == this.stack2[len(this.stack2)-1] {
        this.stack2 = this.stack2[:len(this.stack2)-1]
    }    
}


func (this *MinStack) Top() int {
    return this.stack1[len(this.stack1)-1]
}


func (this *MinStack) Min() int {
    return this.stack2[len(this.stack2)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
```
