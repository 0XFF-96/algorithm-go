package tree

import (
	"fmt"
	"math"
	"testing"
)

// https://www.youtube.com/watch?v=-41E1g8NoMM
// 理解题目意思
// all 【 left subtree node values 】and the sum of 【all right subtree node values】
// 不是简单的两个 nodes 之间的解...
// 这道题目，没有想出来的原因是因为
// case2: 不够明确导致的, 没有真正理解 tilt 的含义。
//
func TestFindTilt(t *testing.T){
	tree1 := initTree()

	sum := findTilt(tree1)
	t.Log(sum)
}



// 不能用中序遍历解法
//
func findTilt(root *TreeNode) int {
	var tilts []int

	var helper func(root *TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return root.Val
		}
		left := helper(root.Left)
		right := helper(root.Right)

		var ret int
		if left > right {
			ret = left - right
		} else {
			ret = right - left
		}
		fmt.Println(tilts)
		fmt.Println(left, right)
		tilts = append(tilts, ret)
		return ret
	}
	helper(root)
	var sum int
	for _, v := range tilts {
		sum += v
	}
	return sum
}


// 简单递归的做法
// 看看答案
// 后向遍历的模板
// https://leetcode.com/problems/binary-tree-tilt/discuss/102334/Java-Solution-post-order-traversal


func findTiltV2(root *TreeNode) int {
	tilt := 0

	traverse(root, &tilt)
	return tilt
}

func traverse(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}

	left := traverse(root.Left, tilt)
	right := traverse(root.Right, tilt)

	*tilt += max(left, right)
	return left + right + root.Val

}

func findtiltHelper(root *TreeNode, ans *int) int {
	if root == nil { return 0 }
	left := findtiltHelper(root.Left, ans)
	right := findtiltHelper(root.Right, ans)
	*ans += int(math.Abs(float64(left) - float64(right)))
	return root.Val + left + right
}
