package array


// https://leetcode.com/problems/maximum-subarray/discuss/20193/DP-solution-and-some-thoughts
// DB 问题的相关解法
// https://leetcode.com/problems/maximum-subarray/
// https://leetcode.com/problems/maximum-subarray/discuss/20193/DP-solution-and-some-thoughts
// 视频
// https://www.youtube.com/watch?v=2MmGzdiKR9Y
// https://www.youtube.com/watch?v=86CQq3pKSUw
//
// 优化问题 dp：动态规划转移方程
// 1、什么解法？
// 2、https://www.youtube.com/watch?v=7J5rs56JBs8
// 3、
// 4、
// wrong answer!
import "math"
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}

	maxSum := math.MinInt32
	curSum := nums[0]

	for i:=1;i < len(nums); i++ {
		tmp := curSum + nums[i]
		if tmp > 0 {
			curSum = tmp
			if tmp > maxSum {
				maxSum = tmp
			}
		} else {
			curSum = 0
		}
	}
	return maxSum
}


// 多分析一下现成的答案
// https://leetcode.com/problems/maximum-subarray/discuss/225619/golang-dp-solution