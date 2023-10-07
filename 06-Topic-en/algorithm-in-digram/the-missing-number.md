# The missing number 

```

func missingNumber(nums []int) int {
    left, right := 0, len(nums) -1

    for left <= right {
        m := left + (right - left) / 2 

        if nums[m] == m {
            left = m + 1 
        } else {
            right = m - 1 
        }
    }

    return left 
}

```
