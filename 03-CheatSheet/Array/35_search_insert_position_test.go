package array

// 代码比较差。。
// 有几个地方的取值点，边界条件需要注意的。
// 能不不能把他们融合在一起？
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}
	idx :=0
	for i :=1;i < len(nums);i++ {
		if nums[i-1] < target && nums[i] >= target {
			idx = i
		}
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target == nums[len(nums)-1] {
		return len(nums)-1
	}
	return idx
}
// case 1: 开头插入
// case 2: 中间插入
// case 3: 相同元素插入
// case 4: 末尾插入