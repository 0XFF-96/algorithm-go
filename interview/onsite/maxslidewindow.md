
- 滑动窗口最大值

```
 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。  
func maxslidewindow(nums []int, k int) []int {
    if len(nums) == 0 {
        return nil
    }
    ret := make([]int, 0, len(nums) - k)
    slide := make([]int, 0, k)
    for _, n := range nums {
        // 如何把元素加入滑动窗口
        // 1. 单调队列
        // 单调队列， 单调递增
        // 三个操作：
        // append
        //   n < slide 的最后一位， 加入
        //   n >= slide 的最后一位，保持单调性
        //   把 slide 删除到最后一位
        // pop
        // move
        if n < slide[len(slide)-1] || len(slide) == 0 {
            slide = append(slide, n)
        } else {
            // 保持单调性
            // 从 slide 数组中，
            // 删除比 n 小的元素
            for slide[len(slide)-1] < n {
                slide = slide[:len(slide)-1]
            }
        }
        // 如果缩小滑动窗口
        // 需要维护 slide 为单调队列
        if len(slide) > 0 && n == slide[0] {
            // 滑动窗口向右移动
            slide = slide[1:]
        }
        // 把滑动窗口的最大值，存入结果数组
        ret = append(ret, slide[0])
    }
    return ret
}
```
