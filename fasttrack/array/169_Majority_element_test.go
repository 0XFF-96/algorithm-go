package array


// 一次 AC
// 使用 Counter-part 的技巧进行
// https://leetcode.com/articles/majority-element/
// 有 5 种方法
// 需要一一实现！
func majorityElement(nums []int) int {
	// Majority Element
	if len(nums) == 0 {
		return 0
	}
	ele := nums[0]
	count := 1
	for i:=1;i < len(nums);i++ {
		if nums[i] == ele {
			count++
		} else {
			count--
		}
		if count == 0 {
			count = 1
			ele = nums[i]
		}
	}

	return ele
}



