类型：优化题、DP、数组
技巧：转化为问题求解
关键点：找到递归公式 maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];  

reference :https://leetcode.com/problems/maximum-subarray/discuss/395100/Go-golang-two-solutions

```golang 
Analysis of this problem:
1-this is a optimization problem, which can be usually solved by DP. 
2- we need to figure out the format of the sub problem. 

3-maxSubArray(int A[], int i, int j) why it is not a good choice ?

> it's hard to find the connection from the sub problem to the original problem(at least for me). In other words, I can't find a way to divided the original problem into the sub problems and use the solutions of the sub problems to somehow create the solution of the original one.


So I change the format of the sub problem into something like: maxSubArray(int A[], int i), which means the maxSubArray for A[0:i ] which must has A[i] as the end element. Note that now the sub problem's format is less flexible and less powerful than the previous one because there's a limitation that A[i] should be contained in that sequence and we have to keep track of each solution of the sub problem to update the global optimal value. However, now the connect between the sub problem & the original one becomes clearer
```

