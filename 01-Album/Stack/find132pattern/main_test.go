



// 错误的做法
// [3,5,0,3,4] 这个方法错误
// 3,5,4 是一种解
func find132pattern(nums []int) bool {
    if len(nums) < 3 {
        return false 
    }
    // 滑动窗口的概念
    //  i < j < k and ai < ak < aj.
    ai := nums[0]
    aj := nums[1]
    
    // [3,5,0,3,4] ?
    // 这个 test case 
    for k := 2; k < len(nums); k++ {
        if ai < nums[k] && nums[k] < aj {
            return true 
        }
        ai = aj 
        aj = nums[k]
    }
    return false 
}