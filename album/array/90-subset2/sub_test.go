package _0_subset2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	// 排序也很关键，
	// 因为 [4, 4, 4, 1, 4], 这种情况
	// 一行代码
	sort.Ints(nums)
	helper(nums, 0, []int{}, &ans)
	return ans
}

func helper(nums []int, idx int, set []int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, set...))
	//fmt.Println("------", ans, idx, set)
	for i:=idx; i<len(nums); i++ {
		// 两行代码
		// 能否画出程序的递归流程图？
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		set = append(set, nums[i])
		// 为什么需要加 1 ？
		helper(nums, i+1, set, ans)
		set = set[:len(set)-1]
	}
}
