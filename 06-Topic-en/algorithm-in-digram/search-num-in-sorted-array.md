# Search numbers in sorted array 

```
func search(nums []int, target int) int {
    // 二分搜索，然后定位左右边界的相关问题。 
    // 如果先判断， 然后移动 左右指针 ？

    // right boundary - left boundary 
    return searchBoundary(nums, target) - searchBoundary(nums, target - 1) 
}

func searchBoundary(nums []int, target int) int {
    i, j := 0, len(nums) - 1 
    for i <= j {
        m := ( i + j ) / 2

        if nums[m] <= target { // 小于等于
            i = m + 1 // 比较重要
        } else {
            j = m - 1 
        }
    }
    return i
}

```
