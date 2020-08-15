package main 
// 两指针的方法
// 
func moveZeroes(nums []int)  {
    if len(nums) <= 1 {
        return 
    }
    l:=0
    for idx, v := range nums {
        if v != 0 {
            nums[idx], nums[l] = nums[l], nums[idx]
            l++
        }
    }
    return 
}
