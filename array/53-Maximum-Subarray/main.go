package main

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = Max(nums[i], dp[i-1]+nums[i])
	}
    return Max(dp...)
}

func Max(slice ...int) int {
	var m int
	for i, v := range slice {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}