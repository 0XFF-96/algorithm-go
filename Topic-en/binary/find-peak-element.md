
# find peak element 

```

func findPeakElement(nums []int) int {
    n := len(nums)

    // 辅助函数，输入下标 i，返回 nums[i] 的值
    // 方便处理 nums[-1] 以及 nums[n] 的边界情况
    get := func(i int) int {
        if i == -1 || i == n {
            return math.MinInt64
        }
        return nums[i]
    }

    left, right := 0, n-1
    for {
        mid := (left + right) / 2
        if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
            return mid
        }
        if get(mid) < get(mid+1) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
}

```


# 最简单的方法， 相当于滑动窗口

```
func peakIndexInMountainArray(A []int) int {

	for i, v := range A {
		if i == 0 {
			continue
		} else if v > A[i-1] && v > A[i+1] {
			return i
		}
	}
	return 0
}
```
