package tree

import (
	"fmt"
	"math"
	"testing"
)

// 思考漏洞的一个点
// 1、没有思考以前是不是有题目类似于这题 path sum
// 2、用递归的方法解题

func TestMaxPathSum(t *testing.T){
	tree1 := initTree()

	a := maxPathSum(tree1)
	t.Log(a)
}


func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt32

	var helper func(root *TreeNode) int
	helper = func (root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := max(0, helper(root.Left))
		r := max(0, helper(root.Right))

		fmt.Println(root.Val)
		sum := l + r + root.Val
		ans = max(ans, sum)
		return max(l, r) + root.Val
	}

	helper(root)

	return ans
}



// 错误解法
//
func pathSumVV(root *TreeNode) int {
	if root == nil || root.Val < 0 {
		return 0
	}

	left := pathSumVV(root.Left)
	right := pathSumVV(root.Right)

	// 负数的情况下
	if left > right {
		return root.Val + left
	} else {
		return root.Val + right
	}

}

// 解法2
// 同样用递归的方式去解决
// 但是维护全局变量的方式有所不同


// 其他两个不同的视频：https://www.youtube.com/watch?v=mOdetMWwtoI
// https://www.youtube.com/watch?v=cSnETAcziS0