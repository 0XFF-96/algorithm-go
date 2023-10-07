
- 颜色排序

```

func sortColors(nums []int)  {
    if len(nums) < 2 {
        return 
    }

    p0 := -1 
    i := 0 
    p2 := len(nums)

    for i < p2 {
        if nums[i] == 0 {
            p0++
            swap(nums, i, p0)
            i++
        } else if (nums[i] == 1) {
            i++
        } else {
            p2--
            swap(nums, i, p2)
        }
    }
}

```
