
- 长度最小的子数组。（套模板题目）

```

func minSubArrayLen(target int, nums []int) int {
    if len(nums) == 0 {
        return 0 
    }

    sum := 0 
    ans := 1000000000
    left := 0 
    right := 0 
    for right < len(nums) {
        sum += nums[right] // 先增加
        for sum >= target {
            ans = min(ans, right - left + 1)
            sum -= nums[left]
            left++
        }
        right++
    }

    if ans == 1000000000 {
        return 0 
    } else {
        return ans 
    }
}

```

- 移除元素

```
func removeElement(nums []int, val int) int {
    slow := 0 
    for fast := 0; fast < len(nums); fast++ {
        if  nums[fast] != val {
            nums[slow] = nums[fast]
             slow += 1 
        }
    }
    return slow 
}
```
