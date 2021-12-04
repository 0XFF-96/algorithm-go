package main

import (
	"fmt"
)

func main()  {
	ret := sortedSquares([]int{1, 2, 3})
	fmt.Printf("%+v \n",ret)
}

// 思考过程
// https://leetcode-cn.com/problems/squares-of-a-sorted-array/submissions/
func sortedSquares(nums []int) []int {
	// 不会用排序。O(nlogn)
	// 可以利用空间复杂度。
	// 合并两个有序数组的方式
	// 1. negetive slice
	// 2. positive slice
	// 3. merge
	// 算法的时间和空间复杂度都是 O(N), 需要额外的数组
	// 有用到题目的条件。 3. 非递减顺序

	// 优化方法
	// 双指针。解法。

	if len(nums) == 0 {
		return nums
	}

	low := 0
	high := len(nums) - 1
	// ret := make([]int, len(nums)-1)
	ret := make([]int, 0, len(nums)-1)
	for _, _ = range nums {
		ret = append(ret, 0)
	}
	curIdx := high
	// 能否等于
	// case1: [-5, 0, 1]
	// case2: [-1, 0 ,  5]
	// case3: [0, 0, 1, 2]
	// case4: [-3, -2, -1]
	for low <= high  && curIdx >= 0 {
		ll := nums[low] * nums[low]
		hh := nums[high] * nums[high]

		if ll > hh {
			ret[curIdx] = ll
			low++
		} else {
			ret[curIdx] = hh
			high--
		}
		curIdx--
	}

	return ret
}
