### 常见位操作

资源
- http://graphics.stanford.edu/~seander/bithacks.html#ReverseParallel 。 
- 


### 计算汉明距离

```
int hammingWeight(int n) {
    int res = 0;
    while (n != 0) {
        n = n & (n - 1);
        res++;
    }
    return res;
}
```


### 是否 2 的指数
为什么？
性质？
```
boolean isPowerOfTwo(int n) {
    if (n <= 0) return false;
    return (n & (n - 1)) == 0;
}
```


### 查找只出现一次的元素？ 
- 题目泛化，查找出现 N 次的元素？

```
int singleNumber(int[] nums) {
    int res = 0;
    for (int n : nums) {
        res ^= n;
    }
    return res;
}
```