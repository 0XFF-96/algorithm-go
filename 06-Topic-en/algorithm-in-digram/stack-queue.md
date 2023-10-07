
- 用两个栈模拟队列

```

type CQueue struct {
    inStack []int 
    outStack []int
}

func Constructor() CQueue {
    return CQueue{
        inStack: make([]int, 0),
        outStack: make([]int, 0),
    }
}


func (this *CQueue) AppendTail(value int)  {
    this.inStack = append(this.inStack, value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.inStack) == 0 && len(this.outStack) == 0 {
        return -1
    }

    if len(this.outStack) != 0 {
        // pop from the stack 
        popVal := this.outStack[len(this.outStack)-1]
        this.outStack = this.outStack[:len(this.outStack)-1]
        return popVal
    } else {
        // inStack -> outstack , reverse push into outstack 
        for i := len(this.inStack) -1; i >=0; i-- {
            this.outStack = append(this.outStack, this.inStack[i])
        }
        this.inStack = make([]int, 0)
        return this.DeleteHead()
    }
}

```
