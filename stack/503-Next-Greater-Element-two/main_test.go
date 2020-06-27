package _03_Next_Greater_Element_two


func nextGreaterElements(nums []int) []int {
	// 1. search circularity to find its next greater value
	// 2. circular array
	// output - 1 for this number.

	// better brute force
	// var res []int
	res := make([]int, len(nums))
	for i:= 0; i < len(nums); i++{
		res[i] = -1
		for j := 1; j < len(nums); j++ {
			if nums[(i + j ) % len(nums)] > nums[i] {
				res[i] = nums[ (i + j) % len(nums)]
				break
			}
		}
	}
	return res
}

