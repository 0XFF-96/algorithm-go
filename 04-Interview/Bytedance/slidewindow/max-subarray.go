

// 最大自数组和
// maxSubArray() 

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 子数组 是数组中的一个连续部分

// 思维一：
// 1.需要以动态规划的角度去看待这个问题。
// 2.代码很简单，但是得到代码的结论不简单。
// 3. dp[i]：表示以 nums[i] 结尾 的 连续 子数组的最大和。 （定义问题，非常重要） 
// 4. 假设数组 nums 的值全都严格大于 00，那么一定有 dp[i] = dp[i - 1] + nums[i]
// 

// https://leetcode-cn.com/problems/maximum-subarray/solution/dong-tai-gui-hua-fen-zhi-fa-python-dai-ma-java-dai/

// 思维二：线段树的思想（不会)
// https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/

// 思维三： 分治的思想

// 二叉树的最大路径和
func maxSubArray(nums []int) int {
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] + nums[i-1] > nums[i] {
            nums[i] += nums[i-1]
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}