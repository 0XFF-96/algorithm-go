
# 最大元素， 摩尔投票法

```

func majorityElement(nums []int) (ans []int) {
    cnt := map[int]int{}
    for _, v := range nums {
        cnt[v]++
    }
    for v, c := range cnt {
        if c > len(nums)/3 {
            ans = append(ans, v)
        }
    }
    return
}


```


```

// func majorityElement(nums []int) []int {
//     // slide window
//     // 1/2 抵消的做法
//     // couting 
//     // 1. 
//     // 2. 
//     // 3. 
//     // 摩尔投票法。该算法用于1/2情况，它说：“在任何数组中，出现次数大于该数组长度一半的值只能有一个
//     // 于1/3。可以着么说：“在任何数组中，出现次数大于该数组长度1/3的值最多只有两个

//     // 首先要知道，在任何数组中，出现次数大于该数组长度1/3的值最多只有两个。
//     // 我们把这道题比作一场多方混战，战斗结果一定只有最多两个阵营幸存，其他阵营被歼灭。数组中的数字即代表某士兵所在的阵营。

//     // 我们维护两个潜在幸存阵营A和B。我们遍历数组，如果遇到了属于A或者属于B的士兵，则把士兵加入A或B队伍中，该队伍人数加一。继续遍历。
//     // 因次找到第一多和第二多的元素即可，然后判断第二多的元素个数是否超过1/3就可以

// }
```



