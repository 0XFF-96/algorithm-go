func maxProduct(nums []int) int {
    r := nums[0]
    n := len(nums)
    imax := r
    imin := r
    
    for i:= 1; i < n; i++ {
        if nums[i] < 0 {
            imin, imax = imax, imin
        }
        
        if nums[i] > imax * nums[i] {
            imax = nums[i]
        } else {
            imax = imax * nums[i]
        }
        
        if nums[i] < imin * nums[i] {
            imin = nums[i]
        } else {
            imin = imin * nums[i]
        }
        
        if r < imax {
            r = imax
        }

    }
    return r
}