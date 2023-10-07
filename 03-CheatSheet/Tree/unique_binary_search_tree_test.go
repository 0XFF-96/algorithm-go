package tree

import (
	"testing"
)

// 卡特兰数 = Unique Binary Search Trees
// https://zh.wikipedia.org/wiki/%E5%8D%A1%E5%A1%94%E5%85%B0%E6%95%B0
// 类型：找规律的题目，有数学公式可以推导
// 归类为树的性质类


// 分而治之的思想
// 斐波那契树列的思想
// https://leetcode.com/problems/unique-binary-search-trees/discuss/31707/fantastic-clean-java-dp-solution-with-detail-explaination

// 视频为：https://www.youtube.com/watch?v=YDf982Lb84o
// 公式为： c = (2n)!/(n+1)!*n!


// 视频2：
// https://www.youtube.com/watch?v=GgP75HAvrlY
//
func numTree(n int) int {
	res := 0
	if n == 0 || n == 1{
		return 1
	}
	if n == 2 {
		return 2
	}
	for i := 1; i <= n; i++{
		res += numTree(i-1) * numTree(n-i)
	}
	return res
}

// DB solution
func numTreeDP(n int) int {
	c := make([]int, n+1)
	c[0] = 1
	c[1] = 1
	for i := 2; i <= n; i ++{
		for j:=1; j <= i;j++{
			c[i] += c[j-1] * c[i-j]
		}
	}
	return c[n]
}

func TestNumTree(t *testing.T){
	n1 := numTree(3)
	n2 := numTreeDP(3)
	t.Log(n1, n2)
}

func TestNumTrees(t *testing.T){

	// n := numTrees(3)

	n := numTreesRe(3)

	// 相差一个阶的计算方式
	n2 := numTreesWrong(3)
	t.Log(n, n2)
}

// 这样算
// 正好相差一个，阶位置
//
func numTreesWrong(n int) int {
	T := make([]int, n+1)
	T[0] = 1
	T[1] = 1
	for i:=2; i<= n; i++{
		for j:=0;j<i;j++{

			// 这行公式有问题
			//
			// 难点在于理解： T[i] += T[j]*T[i-j]
			T[i] += T[j]*T[i-j-1]
		}
	}
	return T[n]
}



func numTreesRe(n int) int {//总共有区间长度0到区间长度n的n+1种情况，这题先把递归的写出来再推出非递归可能会好理解很多，递归代码我也写在下面了
	xx:=make([]int,n+1)
	xx[0]=1
	xx[1]=1

	// 这个公式推导到问题
	//
	for j:=2;j<n+1;j++{
		for i:=1;i<j+1;i++{
			xx[j]=xx[j]+xx[i-1]*xx[j-i]
		}
	}
	return xx[n]
}

// TODO:不会 generates bianry Tree
// https://leetcode.com/problems/unique-binary-search-trees-ii/discuss/346985/Golang-clean-code
// 95. Unique Binary Search Trees II
// 暴力解法？没有什么


// 另一种解法
// https://leetcode.com/problems/unique-binary-search-trees/discuss/149266/golang-solution
//

// generates bianry tree 另一种解法
// https://www.youtube.com/watch?v=GZ0qvkTAjmw

// generates binary tree 的解法
// https://www.youtube.com/watch?v=GZ0qvkTAjmw
