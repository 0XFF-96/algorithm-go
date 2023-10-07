- 只出现一次的数字。 

```

func singleNumber(nums []int) int {
    single := 0 
    for _, i := range nums {
        single ^= i 
    }
    return single
}

```


```

func singleNumber(nums []int) int {
    one := 0 
    two := 0 

    for _, i := range nums {
        one = one ^ i &(-two)
        two = two ^ i &(-one)
    }
    return one
}

```
