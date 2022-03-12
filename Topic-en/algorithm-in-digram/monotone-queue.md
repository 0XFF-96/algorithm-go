
- monotone queue 
- 单调函数

```

type MaxQueue struct {
    queue []int 
    monoQueue []int // desc order 
}


func Constructor() MaxQueue {
    return MaxQueue{
        queue: make([]int, 0),
        monoQueue: make([]int, 0),
    }
}

func (this *MaxQueue) Max_value() int {
    if len(this.monoQueue) > 0 {
        return this.monoQueue[0]
    } else {
       return -1
    }
}


func (this *MaxQueue) Push_back(value int)  {
    this.queue = append(this.queue, value)
    for len(this.monoQueue) > 0 && this.monoQueue[len(this.monoQueue)-1] < value {
        this.monoQueue = this.monoQueue[:len(this.monoQueue)-1]
    }
    this.monoQueue = append(this.monoQueue, value)
    return 
}


func (this *MaxQueue) Pop_front() int {
    if len(this.queue) == 0 {
        return -1 
    }
    val := this.queue[0]
    this.queue = this.queue[1:]
    for len(this.monoQueue) > 0  && this.monoQueue[0] == val {
        this.monoQueue = this.monoQueue[1:] // cut
    }
    return val 
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

 ```
