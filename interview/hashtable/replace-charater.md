
- 错误解法，但是能够跑通。 
- 错误原因是，没有看到题目的返回要求和限制条件。 

```
func characterReplacement(s string, k int) int {
    // 计算最长重复字符 maxRepeatCount
    // 计算第二长的重复字符 secondMaxRepeatCount
 
    // 第二长的重复字符 和 k 的关系 
    // return maxRepeatCount + min(secondMaxRepeatCount, k)
    set := map[rune]int{}
    for _, ss := range s {
        set[ss] += 1 
    }

    // 如果有很多个 maxRepeat 怎么办？
    maxRepeatCount := 0 
    secondMaxRepeatCount := 0 
    for _, i := range set {
        if i > secondMaxRepeatCount && (i < maxRepeatCount || maxRepeatCount == 0 ){
            secondMaxRepeatCount = i 
        }

        // 如果第一个是第二大的呢？
        if i > maxRepeatCount {
            maxRepeatCount = i 
        }
    }
    // 【包含相同字母】 的【最长】 【子字符串】 的长度
    return maxRepeatCount + min(secondMaxRepeatCount, k)
}
```

- 最大连续 1 的个数 III
- 滑动窗口 + 哈希表的相关解题方法

```
func longestOnes(nums []int, k int) int {
    // 翻转 0 的意思是， 把 0 变为 1 。 
    maxOne := 0 
    left, right := 0, 0 
    for right < len(nums) {
        // 向右移动窗口
        if nums[right] == 1 {
            maxOne += 1 
        }

        if right - left + 1 > maxOne + k { // 重要判断条件
            if nums[left] == 1 {
                maxOne -= 1 
            }
            left += 1
        }

        // 窗口一直向右移动
        right += 1 
    }
    return right - left 
}

```
