
### 两个数组的交集

- 成功添加后，需要移除元素，否则会产生重复。 


```
func intersection(nums1 []int, nums2 []int) []int {
    // Why is this tagged as the binary search?
    
    set := map[int]struct{}{}
    for _, i := range nums1 {
        set[i] = struct{}{}
    }

    res := make([]int, 0)
    for _, n := range nums2 {
        if _, ok := set[n]; ok {
            res = append(res, n)
            delete(set, n) // 不再需要了，避免重复
        }
    }
    return res 
}
```

- 和上面题目不一样的是， 这里的要求。 

https://www.youtube.com/watch?v=Bp7fGofslng.

-  Each element in the result must appear as many times as it shows in both arrays 
-  And you may return the result in any order.

```

func intersect(nums1 []int, nums2 []int) []int {
    // Why is this tagged as the binary search?
    
    // 怎么显示次数？
    //  Each element in the result must appear as many times as it shows in both arrays 
    set := map[int]int{} // element:time
    for _, i := range nums1 {
        set[i] += 1 
    }

    res := make([]int, 0)
    for _, n := range nums2 {
        if i, ok := set[n]; ok && i > 0 {
            res = append(res, n)
            set[n] -= 1 
            // delete(set, n) // 不再需要了，避免重复
        }
    }
    return res 
}

```
