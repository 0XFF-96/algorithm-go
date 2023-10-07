// 滑动窗口最大值

```
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return []int{}
    }

    queue := make([]int, 0)

    // len(nums) - k + 1 ， 为什么加 1 ？
    // ret := make([]int, 0, len(nums) - k + 1)
    ret := make([]int, len(nums) - k + 1 )
    left := 1 - k 
    right := 0 
    for right < len(nums) {

        if left > 0 && queue[0] == nums[left - 1] {
            // remove first item from queue 
            // 收缩窗口
            queue = queue[1:]
        }

        for len(queue) > 0 && queue[len(queue)-1] < nums[right] {
            // 维护单调队列特性
            // 移除最后一个元素
            queue = queue[:len(queue)-1]
        }
        // add 
        queue = append(queue, nums[right])
        if left >= 0 {
            ret[left] = queue[0]
            // ret = append(ret, queue[0]) // 初始化方法同，导致这里的代码不一样。
        }
        left++
        right++
    }
    return ret 
}
```
