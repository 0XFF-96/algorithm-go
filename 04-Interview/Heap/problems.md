- heapify, 利用 golang 内置接口，进行最小堆构建。 

```
func getLeastNumbers(arr []int, k int) []int {
    // 构建最小堆
h:=&MyHeap{}
heap.Init(h)
for j:=0;j<len(arr);j++{
    heap.Push(h,arr[j])
}
var result []int
    //最小堆 堆顶元素Pop k次，并存入结果数组中
    for j:=0;j<k;j++{
        min:=heap.Pop(h)
        result=append(result,min.(int))
    }
    return result
}
// golang 最小堆
type MyHeap []int

func (h *MyHeap) Len()int{
    return len(*h)
}
func (h *MyHeap)  Less(i,j int)bool{
    return (*h)[i]< (*h)[j]// 较小的元素会被移动到堆顶
}
func (h *MyHeap)Swap(i,j int){
    (*h)[i],(*h)[j]=(*h)[j],(*h)[i]
}
func (h *MyHeap) Push(x interface{}){
    (*h)=append((*h),x.(int))
}
func (h *MyHeap) Pop()interface{}{

    v:=(*h)[h.Len()-1]
    *h=(*h)[:h.Len()-1]
    return v
}
```
