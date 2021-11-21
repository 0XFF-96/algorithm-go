package array


// two-point, in-place, 元素交换题目
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i :=0
	for j:=0;j<len(nums);j++{
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 还有没有其他方式解法？
// https://leetcode.com/problems/remove-element/discuss/12420/Golang-concise-solution
// 这种方法更 intuitive
func removeElementV2(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j:=0; j < len(nums);j++{
		// 用这个判断条件可以吗？
		if nums[j] == val {

		}
	}
	return i
}
