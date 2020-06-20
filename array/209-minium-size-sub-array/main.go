func minSubArrayLen(s int, nums []int) int {
    // contiguous subarray 
    // 动态规划
    // O(N) 和 O(n log n) 算法
    // 暴力解法
    // for for 循环
    // prefix array 
    
    // 不过： [1,4,4]
    ret := 0
    for i :=0; i < len(nums); i ++ {
        curSum := nums[i]
        if curSum == s {
            return 1
        }
        k := 0
        for j:=i+1; j < len(nums);j++{
            if curSum >= s && (k < ret || ret == 0){
                ret = k
                break
            }
            curSum += nums[j]
            k++
        }
        // 1, 4, 4 
        // 
        // if curSum >= s && ( k < ret || ret == 0 ){
        //     ret = k
        // }
    }
    return ret
}


// solution 2 
// 能够解决
// 为什么上面的不能够解决？
// 不能够 AC ?

// 优化点1： k 变量
// 优化点2： for 循环♻️

func minSubArrayLen(s int, nums []int) int {
    // contiguous subarray 
    // 动态规划
    // O(N) 和 O(n log n) 算法
    // 暴力解法
    // for for 循环
    // prefix array 
    
    // 不过： [1,4,4]
    ret := 0
    for i :=0; i < len(nums); i ++ {
        curSum := 0
        for j:=i; j < len(nums);j++{
            curSum += nums[j]
            if curSum >= s && ( j-i + 1 < ret || ret == 0){
                ret = j-i +1
                break
            }
        }
    }
    return ret
}

// https://leetcode.com/problems/minimum-size-subarray-sum/solution/

