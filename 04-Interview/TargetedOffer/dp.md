

- dp 数组

```
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0 
    }
    res := nums[0]
    for i:=1; i < len(nums); i++{
        if nums[i-1] > 0 {
            nums[i] += nums[i -1]
        }
        res = max(res, nums[i])
    }
    return res
}

func max(i, j int) int {
    if i > j {
        return i
    } else {
        return j
    }
}
```


- 礼物的最大价值


```
func maxValue(grid [][]int) int {
    row := len(grid)
    column := len(grid[0])

    for i := 0; i < row; i++{
        for j := 0; j < column; j++{
            if i == 0 && j == 0 {
                continue 
            }
            if i == 0 && j != 0 {
                // 只能向右
                grid[i][j] += grid[i][j-1]
            }

            if i != 0 && j == 0 {
                // 只能向下
                grid[i][j] += grid[i-1][j]
            }

            if i!= 0 && j != 0 {
                grid[i][j] += max(grid[i][j-1], grid[i-1][j])
            }
        }
    }

    // 因为边界条件是 [), 左开右闭， 判断条件是 < 
    // 所有这里是需要用 - 1 
    return grid[row -1][column - 1]
}

func max(i , j int) int {
    if i > j {
        return i
    } else {
        return j 
    }
 
```


- 股票最大利润


```
func maxProfit(prices []int) int {
    minCost := 1000000000 
    maxProfit := 0 
    for _, p := range prices {
        minCost = min(minCost, p)
        maxProfit = max(maxProfit, p - minCost)
    }
    return maxProfit
}
```


