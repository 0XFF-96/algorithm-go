package _0_remove_duplicate_2



// 没有 AC
// 对于 left 和 count 的控制不好
func removeDuplicates(nums []int) int {
	// 双指针的解题方法
	left := 0
	count := 0
	for i := 0; i < len(nums); i ++{
		if i != 0 && nums[left] == nums[i] {
			if count == 2 {
				continue
			} else {
				left++
				if left <= i {
					nums[left] = nums[i]
					count++
				}
			}
		} else {
			// 赋值
			// left++
			nums[left] = nums[i]
			count = 0
		}
		count++
	}

	nums = nums[:left+1]
	return left +1
}


// 去除多余变量 count
// 代码会清晰非常多！
func removeDuplicatesV2(nums []int) int {
	// 双指针的解题方法
	j := 0
	for i :=0; i < len(nums); i ++{
		// nums[i] != nums[i-2] !!!
		if j < 2 || nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j ++
		}
	}
	return j
}
