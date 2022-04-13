### Exchange 


```

func exchange(nums []int) []int {
    i, j := 0, len(nums) -1 

    for i < j {
        for i < j && nums[i] & 1 == 1 {
            i += 1 
        }

        for i < j && nums[j] & 1 == 0 {
            j -= 1 
        }

        // 否则交换
        nums[i], nums[j] = nums[j], nums[i]
    }
    return nums
}

```
