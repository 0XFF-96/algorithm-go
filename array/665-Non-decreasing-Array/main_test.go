package main 




// 不 AC 的方案
// 没有考虑边界条件
// 需要考虑各种输入
// [4,2,3]

// [3,4,2,3]
func checkPossibilityV1(nums []int) bool {
    if len(nums) <= 1 {
        return true 
    }
    // 排序的技巧
    // 查找逆序对
    c := 0 
    for i:=1; i < len(nums); i ++ {
        if nums[i-1] > nums[i] {
            c++
        }
        if c >= 2 {
            return false
        }
    }
    return true 
}



// 
func checkPossibility(nums []int) bool {
    count := 0
    index := 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            index = i
            count++
        }
    }
    if count == 0 { return true }
    if count == 1 {
		if index == 1 || index == len(nums)-1 { return true }
		
		// 这里是为了
		// [3,4,2,3]
		// 这种情况
        if nums[index+1] >= nums[index-1] || nums[index] >= nums[index-2] { return true }
    }
    return false
}