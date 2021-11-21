package array


func findLengthOfLCIS(nums []int) int {
	ans := 0
	anchor := 0

	for i:=0;i < len(nums);i++{
		if i > 0 && nums[i-1] >= nums[i] {
			anchor = i
		}
		ans = max(ans, i - anchor + 1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 另一种解法
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/discuss/107365/JavaC%2B%2BClean-solution
//



