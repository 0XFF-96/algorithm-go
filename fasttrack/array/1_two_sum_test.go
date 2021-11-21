package array


// 需要写出三种解法
// 优化的思想：固定某个变动参数。

// one-pass solution
func twoSum(nums []int, target int) []int {
	if len(nums) ==0 {
		return nil
	}
	vMap := map[int]int{}
	for i:=0;i < len(nums);i++{
		complement := target - nums[i]
		if idx, ok := vMap[complement];ok{
			return []int{idx,i}
		}
		vMap[nums[i]] = i
	}
	return nil
}
