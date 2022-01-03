### 题目类型
1. https://www.geeksforgeeks.org/solve-dynamic-programming-problem/


2. 

### NOTE
1. 这些题目的变种很多，需要了解每一个条件背后代表的含义是否清晰。
2. 
条件2:
```
Design an algorithm to find the maximum profit. You may complete as many transactions as you like
```

### Note 
```
Time complexity : O(n)O(n). Only a single pass is needed.

Space complexity : O(1)O(1). Only two variables are used.

初始想法

Time complexity : O(n)O(n). Only a single pass is needed.
func maxProfit(prices []int) int {
    // 双指针
    // [7,6,4,3,1]
    // [7,6,4,3,1]
    // 了解的信息
    // 知道当前 index 之后的最大值是什么。
    // 懂了, 从后往前扫描一遍，后缀最大值数组
    // O(N) O(N) 时间
}

```

```
更好的做法
func maxProfit(prices []int) int {
    // 双指针
    // [7,6,4,3,1]
    // [7,6,4,3,1]
    // 了解的信息
    // 知道当前 index 之后的最大值是什么。
    // 懂了, 从后往前扫描一遍，后缀最大值数组
    // O(N) O(N) 时间
}
```


```
brute force

public class Solution {
    public int maxProfit(int prices[]) {
        int maxprofit = 0;
        for (int i = 0; i < prices.length - 1; i++) {
            for (int j = i + 1; j < prices.length; j++) {
                int profit = prices[j] - prices[i];
                if (profit > maxprofit)
                    maxprofit = profit;
            }
        }
        return maxprofit;
    }
}
```

func maxProfit(prices []int) int {
    // not engaing multiple transactions at the same time
    // 只要后面有收益，就 buy
    // 简单的二维搜索问题
    // for-for 循环，暴力点就
    // [7, 1, 5, 3, 6, 4]
    // [1, 2, 3, 4, 5]
    // Time complexity : O(n^n)O(n 
    // n
    // ). Recursive function is called n^2
    // Space complexity : O(n)O(n). Depth of recursion is n
    if len(prices) == 0 {
        return 0 
    }
    maxProfit := 0
    for i:= 1; i < len(prices); i ++ {
        if  prices[i] > prices[i-1] {
            maxProfit += prices[i] - prices[i-1]
        }
    }
    return maxProfit
}

### solution NOTE2 
> 这个需要好好研究一下

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/solution/

// Approach 2: Peak Valley Approach
