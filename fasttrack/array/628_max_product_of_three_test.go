package array
import "sort"
func maximumProduct(nums []int) int {
	sort.Ints(nums)

	n1 := nums[0]*nums[1]*nums[len(nums)-1]
	n2 := nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3]
	return max(n1, n2)
}

// 有一个 one-pass 的方案
// 不