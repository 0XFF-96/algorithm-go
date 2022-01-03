func productExceptSelf(nums []int) []int {
    // solve it witout division 
    // 前缀和？
    // 后缀和？
    //  It's guaranteed that the product of the elements of any prefix or suffix of the array 
    // (including the whole array) fits in a 32 bit integer
    
    // case 
    // 为什么这种方法能够行。
    // 1，  2，  3， 4 
    // 24， 12， 4， 1 
    // 1，  1，  2， 6 
    // 使用前后两个数组
    // 分别记录，前后两个相乘的积
    
    // 怎么优化空间复杂度
    // 位运算？
    prefix := make([]int, len(nums))
    suffix := make([]int, len(nums))
    
    for idx, _ := range nums {
        if idx == 0 {
            prefix[idx] = 1 
        } else {
            // prefix = append(prefix, nums[idx-1] * prefix[idx-1])
            prefix[idx] = nums[idx-1] * prefix[idx-1]
        }
    }
    
    for i := len(nums)-1; i >=0; i-- {
        if i == len(nums)-1 {
            suffix[i] = 1
        } else {
            // suffix = append(suffix, nums[idx+1] * suffix[idx+1]) 
            suffix[i] = suffix[i+1] * nums[i+1] 
        }
    }
    
    ret := []int{}
    for i:=0; i < len(nums); i ++ {
        ret = append(ret, prefix[i] * suffix[i])
    }
    return ret 
}