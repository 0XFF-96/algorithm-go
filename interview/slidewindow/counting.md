- counting , 

```
func numSubarrayBoundedMax(nums []int, left int, right int) int {
    // 这里为啥不用加 1， 而是需要 left - 1 
    return lessEqualsThan(nums, right) - lessEqualsThan(nums, left - 1)
}

func lessEqualsThan(nums []int, k int) int {
    length := len(nums)
    res := 0 
    left, right := 0, 0 
    for right < length {
        if nums[right] > k {
            left = right + 1 
        }
        res += right - left 
        right++
    }
    return res 
}

```

- 例题：至多包含两个不同字符的最长子串







- 水果成蓝

```

func totalFruit(fruits []int) int {
    if len(fruits) == 0 {
        return 0 
    }

    ans := 0
    cnt := 0 
    left, right := 0, 0 
    set := map[int]int{}
    for right < len(fruits) { 

        if set[fruits[right]] == 0 { // 新加入的
            cnt++
        }
        set[fruits[right]]++
        right++ // 指针右移动
        for cnt > 2 {
            // 窗口收缩
            set[fruits[left]]--
            if set[fruits[left]] == 0 {
                cnt--
            }
            left++
        }
        
        // 要的是，最多可以采集多少颗树
        ans = max(ans, right - left) // 为什么不用 + 1 , 因为 right 是能够取到的
    }

    return ans 
}

```
