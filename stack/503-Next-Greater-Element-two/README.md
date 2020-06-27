### 相关题目
1. https://www.youtube.com/watch?v=5IAq5jd8O7Y
2. 



### 技巧与方法
1. Approach #3 Using Stack [Accepted]
2. brute force 方法。
3. 怎么把一个数组变为循环数组呢？



### 将数组变为循环数组的方法

将： [1 2 1]
变为 [1 2 1 1 2 1]
结果: 可以循环搜索 idx 前面的元素。 
next greater element !

通过这一行语句实现：
```
nums[(i + j ) % len(nums)] > nums[i]
```
