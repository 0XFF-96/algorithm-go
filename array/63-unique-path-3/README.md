
Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths II.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Unique Paths II.
Next challenges:


这个题目相比上一道题目有三个坑：
> 这三个坑分别对应的测试用例是 xxx
[[1]] [[0]]
[[0,0],[1,1],[0,0]]


1- 	obstacleGrid[0][0] = 1
2- 	for i:=1; i < row; i ++ {
3- obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1 ] == 1 
4- obstacleGrid[row -1][col-1]