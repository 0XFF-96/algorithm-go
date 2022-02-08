
- rob house

```
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }

    dp0 := nums[0]
    dp1 := max(dp0, nums[1])
    ans := max(dp0, dp1)

    for i := 2; i < n; i++{

        // i 时刻的状态
        // 只与， i -1, i-2 时刻的状态有关。
        curI := max(dp1, dp0 + nums[i])

        // 滚动数组，
        // 用于压缩状态
        dp0 = dp1 
        dp1 = curI 

        // 是否应该更新最大值
        ans = max(ans, curI)
    }
    return ans 
}
```
