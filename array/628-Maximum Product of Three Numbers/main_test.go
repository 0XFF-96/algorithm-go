
import "sort"
func maximumProduct(nums []int) int {
    
    sort.Ints(nums)
    
    n1 := nums[0]*nums[1]*nums[len(nums)-1]
    n2 := nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3]
    return max(n1, n2)
}

func max(a, b int) int {
    if a > b {
        return a 
    } else {
        return b 
    }
}



// single scan 
// 找出 min1, min2, min3
// 找出 max1, max2, max3 
// 将上面的数字进行组合。 
// https://leetcode.com/problems/maximum-product-of-three-numbers/solution/