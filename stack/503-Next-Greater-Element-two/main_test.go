package _03_Next_Greater_Element_two

func nextGreaterElements(nums []int) []int {
	// 1. search circularity to find its next greater value
	// 2. circular array
	// output - 1 for this number.

	// better brute force
	// var res []int
	// 这是是有数组循环的。
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

func nextGreaterElementsV2(nums []int) []int {
    // 1. search circularity to find its next greater value 
    // 2. circular array 
    // output - 1 for this number.
    // better brute force 
    // var res []int 
    if len(nums) == 0 {
        return nil 
    }
    // 必须循环两遍数组
    res := make([]int, len(nums))    
    stack := []int{}

    for i := 2 * len(nums); i >=0; i--{
        for len(stack) != 0 && nums[stack[len(stack)-1]] <= nums[i % len(nums)] {
            stack = stack[:len(stack)-1]
        }
        
        if len(stack) == 0 {
            res[i%len(nums)] = -1
        } else {
            res[i%len(nums)] = nums[stack[len(stack)-1]]
        }
        stack = append(stack, i % len(nums))
    }
    return res 
}

