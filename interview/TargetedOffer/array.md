
- 二维数组中的查找
```
func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false 
    }

    // search from the left corner 
    x := 0 
    y := len(matrix) -1  // 和下面的 x 统一写法

    // 搜索区间是？
    // 边界条件的处理
    for x <= len(matrix[0]) -1 && y >= 0 {
        if matrix[y][x] > target {
            y--
        } else if matrix[y][x] < target {
            x++
        } else {
            return true 
        }
    }
    return false 
}
```

- min from rotated array 
- 本质，利用 mid 和 right ， 判断下降序列在数组的哪一侧。 


```
func minArray(numbers []int) int {
    if len(numbers) == 0 {
        return -1 
    }

    left := 0 
    right := len(numbers) - 1 

    // if will break, if left == right 
    for left < right {
        mid := left + (right - left)/2 

        if numbers[mid] < numbers[right] {
            right = mid 
        } else if numbers[mid] > numbers[right] {
            left = mid + 1 // why not left = mid ?
        } else if numbers[mid] == numbers[right]{
            right --     // why not left++ ?
        }
    }
    
    // why it is the same to return numbers[left]
    return numbers[right]
}
```

```
func search(nums []int, target int) int {
    if len(nums) == 0 { 
        return 0
    }
    return getRightMargin(nums, target) - getRightMargin(nums, target-1)
}

func getRightMargin(nums []int, target int) int {
    left := 0 
    right := len(nums) -1 

    // 结束条件 <= ?
    // 为什么不是 < 
    // 向右边，收缩搜索区间
    // 为什么是 <= , 能否改成 < ? 为什么
    for left <= right {
        mid := left + (right - left)/2 
        if nums[mid] <= target {
            left = mid + 1 
        } else {
            right = mid - 1
        }
    }
    return left 
}
```
