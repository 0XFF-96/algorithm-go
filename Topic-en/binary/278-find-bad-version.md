
# 278 find bad version

```
func firstBadVersion(n int) int {
    low := 1 
    high := n
    
    for low < high {
        mid  :=  low + (high  - low)/2
        if  !isBadVersion(mid)  {
             low = mid + 1
        }else {
            high = mid
        }
    }
    return low  
}
```
