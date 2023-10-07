
# Count Element 

```
// func countElements(nums []int) int {
//     //  return the number of elements that
// }


func countElements(nums []int) int {
    sort.Ints(nums) // 耗费时间
    
    // 比小的小
    // 比大的大
    // 元素。
    // both a strictly smaller and a strictly greater element appear
    min, max := nums[0], nums[len(nums) - 1]
    minC, maxC := 0, 0
    if min == max { return 0 }
    for _, v := range nums {
        if v == min { minC++ }
        if v == max { maxC++ }
    }
    return len(nums) - minC - maxC
}
```
