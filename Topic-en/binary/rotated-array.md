
### Find Min 

- in a rotated array. 

```
func findMin(nums []int) int {
    // This problem is similar to Find Minimum in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?


    // When num[mid] == num[hi], we couldn't sure the position of minimum in mid's left or right, so just let upper bound reduce one.

    lo := 0 
    hi := len(nums) - 1
    mid := 0 
    
    for lo < hi {
        mid = lo + (hi - lo) / 2 
        if (nums[mid] > nums[hi]) {
            lo = mid + 1
        } else if (nums[mid] < nums[hi]) {
            hi = mid 
        } else {
            // when num[mid] and num [hi] are the same 
            hi -- 
        }
    }

    // shrink to the left bound. 
    return nums[lo]
}
```
