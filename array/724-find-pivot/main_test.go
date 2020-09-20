package main 


func pivotIndex(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    totalSum := 0 
    leftSum := 0 
    for _, n := range nums {
        totalSum += n
    }
    
    for i:= 0;i < len(nums); i++{
        if (leftSum == totalSum - leftSum - nums[i]) {
            return i
        }
        leftSum += nums[i]
    }
    return -1
}


