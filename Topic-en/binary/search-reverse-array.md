
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

- 这是在排序数组中搜索的情况。 

```

func search(nums []int, target int) bool {
    n := len(nums)
    if n == 0 {
        return false
    }
    if n == 1 {
        return nums[0] == target
    }
    l, r := 0, n-1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] == target {
            return true
        }
        if nums[l] == nums[mid] && nums[mid] == nums[r] {
            l++
            r--
        } else if nums[l] <= nums[mid] {
            if nums[l] <= target && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if nums[mid] < target && target <= nums[n-1] {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return false 
}

```

- 旋转数组，搜索最小值 

```
func findMin(nums []int) int {
    low, high := 0, len(nums) -1 

    // 向左边界迫近的写法， 收敛于左边界
    // low < high, high = pivot, low = pivot + 1 
    for low < high {
        pivot := low + (high - low) / 2 

        if nums[pivot] < nums[high] {
            high = pivot
        } else {
            low = pivot + 1 
        }
    }
    return nums[low]
}

```
