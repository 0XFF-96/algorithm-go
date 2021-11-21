package array

// 双指针法
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i :=0
	for j:=1;j <len(nums);j++{
		if nums[i] != nums[j] {
			i++;
			nums[i] = nums[j]
		}
	}
	return i+1
}

// 还有其他办法吗？
// 能不能再优化一下代码？
// 看其他人的代码...
//