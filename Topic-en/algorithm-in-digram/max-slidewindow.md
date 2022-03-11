
- 最大滑动窗口

```

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) < k || k <= 0 {
        return []int{}
    }

    // 单调队列, 单调递减
    queue := []int{}
    // 未形成窗口
    for i := 0; i < k; i++{
        for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
            queue = queue[:len(queue)-1] // pop 保持递减
        }
        queue = append(queue, nums[i])
    }
    
    res := []int{queue[0]} // 初始化
    // 形成窗口后 
    for i:= k; i < len(nums) ; i++ {
        if len(queue) > 0 && queue[0] == nums[i-k] {
            queue = queue[1:] // pop left , 窗口收缩
        }

        for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
            queue = queue[:len(queue)-1]
        }
        queue = append(queue, nums[i])
        res = append(res, queue[0])
    }
    return res 
}

```
