
# 寻找有重复的相关数 

# Find the duplidate number . 

```
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.

```

```

func findDuplicate(nums []int) int {
    n := len(nums)
    l, r := 1, n - 1
    ans := -1
    for l <= r {
        mid := (l + r) >> 1
        cnt := 0
        for i := 0; i < n; i++ {
            if nums[i] <= mid {
                cnt++
            }
        }
        if cnt <= mid {
            l = mid + 1
        } else {
            r = mid - 1
            ans = mid
        }
    }
    return ans
}

```

