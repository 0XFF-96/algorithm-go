package main

import (
    "fmt"
)

// 没能够写出来
func canJump(nums []int) bool {
    // jump game 
    // dp ?
    // 了解过。
    // greedy 
    // greedy 的限制？🚫
    // 路径记录？
    // dfs 的方法？
    // 是的，可以利用 dfs 的思想
    // curPos, nums,   return ?
    // 如何变化？
    if len(nums) == 0 {
        return false
    }
    res := false
    var dfs func(nums []int, curPos int) 
    dfs = func(nums []int, curPos int) {
        if curPos >= len(nums) -1 {
            res = true 
            return
        }
        
        value := nums[curPos] 
        for i:= curPos; i <= i + value; i++ {
            dfs(nums, i)
        }
    }
   return res
}

func main(){
    res := canJumpV2([]int{2,3,1,1,4})
    fmt.Println(res)
}
// 死循环了。
// 没有找到哪里的原因。
//

// [1]
// [2,0]
// [2,3,1,1,4]
// [3,2,1,0,4]
// 等情况。
func canJumpV2(nums []int) bool {
    res := false
    var dfs func(nums []int, curPos int) 
    dfs = func(nums []int, curPos int) {
        if curPos >= len(nums) -1 {
            res = true 
            return
        }

        fmt.Println(curPos)
        value := nums[curPos] 
        if curPos + value == len(nums) -1 {
            res = true 
            return 
        }

        // 添加写不对
        // i < len(nums) -1 ?
        // i := curPos + 1?
        for i:= curPos; i < len(nums)-1; i++ {
            fmt.Println(i)
            dfs(nums, i+1)
        }
        return 
    }
    dfs(nums, 0)
   return res
}