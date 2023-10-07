

```
func singleNonDuplicate(nums []int) int {
    // Single Element in a Sorted Array
    // 
    n := nums[0]
    for i := 1; i < len(nums); i++ {
        n ^= nums[i] 
    }

    // 自己 ^ 自己 = 0 ， 异或。 
    return n 
}
```
