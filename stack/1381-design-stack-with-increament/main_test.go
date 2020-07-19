package main 

type CustomStack struct {
    values []int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{
        values: make([]int, 0, maxSize),
    }
}


func (this *CustomStack) Push(x int)  {
    if len(this.values) == cap(this.values){
        return
    }
    this.values = append(this.values, x)
}


func (this *CustomStack) Pop() int {
    if len(this.values) == 0 {
        return -1
    }
    r := this.values[len(this.values)-1]
    this.values = this.values[:len(this.values)-1]
    return r
}


func (this *CustomStack) Increment(k int, val int)  {
    if k > len(this.values) {
        k = len(this.values)
    }
    
    for i := 0; i < k;i++{
        this.values[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */