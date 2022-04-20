
# 在旋转数组中搜索

```

func search(nums []int, target int) int {
    lens := len(nums)
    left, right := 0, lens -1 

    for left <= right { // <=, lens-1
        mid := (left + right) / 2 
        if nums[mid] == target {
            return mid 
        } else if nums[left] <= nums[mid] {
            if nums[left] <= target && target <= nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            if nums[mid] <= target && target <= nums[right] {
                // 在右区间
                left = mid + 1 
            } else {
                right = mid - 1
            }
        }
    }
    return -1 
}

```
