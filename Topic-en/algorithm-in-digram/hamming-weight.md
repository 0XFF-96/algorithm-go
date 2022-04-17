
# hamming weight 


```
func hammingWeight(num uint32) int {
    res := 0 
    for num > 0 {
        res += 1 
        num = num & (num -1)
    }
    return res 
}

```

```
func hammingWeight(num uint32) int {
    res := int(0)
    for num > 0 {
        res = res + int((num & uint32(1)))
        num = num >> 1 
    }
    return res 
}
```
