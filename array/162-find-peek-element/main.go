
func findPeakElement(nums []int) int {
    if len(nums) <= 1 {
        return 0 
    }
    
    
    // Follow up: Your solution should be in logarithmic complexity
    for i := 0; i < len(nums) ; i++ {
        if i == 0  {
            if nums[i] > nums[i+1] {
                return i
            } else {
                continue
            }
        } 
        if i == len(nums) -1 && nums[i] > nums[i-1] {
            return i
        } 

        if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
            return i
        }
    }
    return -1
}

// NOTE: https://leetcode.com/problems/find-peak-element/