package _8_Subset

import (
	"testing"
)

func TestSubSet(t *testing.T){
	nums := []int{1, 2}

	res := subsetsV3(nums)
	t.Log(res)
}


func subsetsV3(nums []int) [][]int {
	ans := [][]int{}
	helper(nums, 0, []int{}, &ans)
	return ans
}

func helper(nums []int, idx int, set []int, ans *[][]int) {
	*ans = append(*ans, append([]int{}, set...))
	//fmt.Println("------", ans, idx, set)
	for i:=idx; i<len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		set = append(set, nums[i])
		// 为什么需要加 1 ？
		helper(nums, i+1, set, ans)
		set = set[:len(set)-1]
	}
}
