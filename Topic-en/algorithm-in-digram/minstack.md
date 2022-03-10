# 最小栈

````
type MinStack struct {
    minStack []int 
    stack []int 
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        minStack: make([]int, 0),
        stack   : make([]int, 0),
    }
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if len(this.minStack) == 0 || this.minStack[len(this.minStack) -1 ] >= x {
        this.minStack = append(this.minStack, x)
    }
}


func (this *MinStack) Pop()  {
    val := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]

    if this.minStack[len(this.minStack)-1] == val {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack) -1 ]
}


func (this *MinStack) Min() int {
    return this.minStack[len(this.minStack)-1]
}
```
