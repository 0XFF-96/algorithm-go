

- 多数元素，利用抵消的思想。
- AC once within 5 minutes 

```
func majorityElement(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    curNum := nums[0]
    curCount := 1 
    for i:=1; i < len(nums); i++ {
        if curCount == 0 {
            curNum = nums[i]
            curCount++
            continue 
        }

        if nums[i] == curNum {
            curCount ++ 
        } else {
            curCount --
        }
    }
    return curNum
}
```
