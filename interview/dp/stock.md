

- max profix 股票的最大利润。 


```

func maxProfit(prices []int) int {
    length := len(prices)
    min := prices[0]
    maxDiff := 0 

    for i := 0; i < length; i ++ {
        maxDiff = max(maxDiff, prices[i] - min)
        if prices[i] < min {
            min = prices[i]
        }
    }
    return maxDiff
}

```
