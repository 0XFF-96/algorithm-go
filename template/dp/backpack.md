### 背包问题

### 思路

1. dp[i][w] 的定义如下：对于前 i 个物品，当前背包的容量为 w，
这种情况下可以装的最大价值是 dp[i][w]。

2. 比如说，如果 dp[3][5] = 6，其含义为：对于给定的一系列物品中，
若只对前 3 个物品进行选择，当背包容量为 5 时，最多可以装下的价值为 6

3. base case 就是 dp[0][..] = dp[..][0] = 0

4. w - wt[i-1] 可能小于 0 导致数组索引越界的问题 ？


### 核心代码

```
int[][] dp[N+1][W+1]
dp[0][..] = 0
dp[..][0] = 0

for i in [1..N]:
    for w in [1..W]:
        dp[i][w] = max(
            把物品 i 装进背包,
            不把物品 i 装进背包
        )
return dp[N][W]
```

### 核心思想
如何表达
1、把物品 i , 装进背包
2、不把物品 i, 装进背包

```
for i in [1..N]:
    for w in [1..W]:
        dp[i][w] = max(
            dp[i-1][w],
            dp[i-1][w - wt[i-1]] + val[i-1]
        )
return dp[N][W]
```

### 无限背包问题

### 核心思路

```
int dp[N+1][amount+1]
dp[0][..] = 0
dp[..][0] = 1

for i in [1..N]:
    for j in [1..amount]:
        把物品 i 装进背包,
        不把物品 i 装进背包
return dp[N][amount]
```

### 核心状态转移方程

- 求解的问题是

```
综上就是两种选择，而我们想求的 dp[i][j] 是「共有多少种凑法」，
所以 dp[i][j] 的值应该是以上两种选择的结果之和：。 （所以不需要用 max, 进行选择）。
```

```
for (int i = 1; i <= n; i++) {
    for (int j = 1; j <= amount; j++) {
        if (j - coins[i-1] >= 0)
            dp[i][j] = dp[i - 1][j] 
                     + dp[i][j-coins[i-1]];   
return dp[N][W]
```

### 能够使用无限硬币的情况

```
dp[i][j-coins[i-1]] 可以包含重复使用第 i 个硬币的情况吗？
```

### 完全背包问题核心代码框架

```
int change(int amount, int[] coins) {
    int n = coins.length;
    int[][] dp = int[n + 1][amount + 1];
    // base case
    for (int i = 0; i <= n; i++) 
        dp[i][0] = 1;

    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= amount; j++)
            if (j - coins[i-1] >= 0)
                dp[i][j] = dp[i - 1][j] 
                         + dp[i][j - coins[i-1]];
            else 
                dp[i][j] = dp[i - 1][j];
    }
    return dp[n][amount];
}
```