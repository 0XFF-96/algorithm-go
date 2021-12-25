### 最小路径和

### 函数签名

int minPathSum(int[][] grid);


### 思路

- 二维矩阵中求最优化问题（最大值或者最小值），肯定需要递归 + 备忘录，也就是动态规划技巧
- 



### 核心递归代码

- 这是一个自顶向下的解法。 
- 
```
int dp(int[][] grid, int i, int j) {
    // base case
    if (i == 0 && j == 0) {
        return grid[0][0];
    }
    // 如果索引出界，返回一个很大的值，
    // 保证在取 min 的时候不会被取到
    if (i < 0 || j < 0) {
        return Integer.MAX_VALUE;
    }

    // 左边和上面的最小路径和加上 grid[i][j]
    // 就是到达 (i, j) 的最小路径和
    return Math.min(
            dp(grid, i - 1, j), 
            dp(grid, i, j - 1)
        ) + grid[i][j];
}
```