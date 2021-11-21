
package archive_Tree

// 1、利用中序遍历解法
// Follow-up-Question: 如果不是 BST 怎么办？
// 1、用 TreeSet 记录，寻找 floor and  ceiling of cur.Val
// 难点在哪里？
// 
// 方法2:

// 有几个点可以提示用 中序遍历，1、BST 2、最小 diff 一定存在于两个相邻元素之间
//
// 没有 AC 的答案
// 难点在于：如何初始化第一个 中序遍历的元素？
// 有个 bug 在这里...
import (
	"fmt"
	"math"
	"testing"

	"pratice-go/algorithmTopic/tree"
)
func getMinimumDifference(root *tree.TreeNode) int {
	cur := root
	minDef := math.MaxInt64
	preVal := 0

	// 怎么初始化第一个节点...
	for cur != nil {
		if cur.Left == nil {
			if minDef == math.MaxInt64 {
				preVal = cur.Val
				cur = cur.Right
				continue
			}

			curDiff := MinDiff(cur.Val, preVal)
			if curDiff < minDef {
				minDef = curDiff
			}
			cur = cur.Right
		} else {

			// 中序遍历下的前一个节点。
			// 名字应该怎么叫？
			preNode := cur.Left
			for preNode.Right != nil && preNode.Right != cur {
				preNode = preNode.Right
			}
			if preNode.Right == nil {
				preNode.Right = cur
				cur = cur.Left
			} else {
				preNode.Right = nil

				if minDef == math.MaxInt64 {
					preVal = cur.Val
				}

				curDiff := MinDiff(cur.Val, preVal)
				if curDiff < minDef{
					minDef = curDiff
				}

				cur = cur.Right
			}

		}
	}

	return minDef
}

func TestGetMin(t *testing.T){
	tree1 := &tree.TreeNode{Val: 1}

	tree1.Right = &tree.TreeNode{Val: 3}
	tree1.Right.Right = &tree.TreeNode{Val: 2}

	min := getMinimumDifferenceV2(tree1)
	t.Log(min)

}

func getMinimumDifferenceV2(root *tree.TreeNode) int {
	minDef := math.MaxInt64
	preVal := math.MinInt32

	var helper = func(root *tree.TreeNode){}

	helper = func(root *tree.TreeNode){
		if root == nil {
			return
		}
		helper(root.Left)
		if preVal != math.MinInt32 {
			curDef := MinDiff(root.Val, preVal)
			if curDef < minDef {
				minDef = curDef
			}
			fmt.Println("curDef", curDef, preVal, root.Val)
		}
		preVal = root.Val
		fmt.Println("traver:", root.Val)
		helper(root.Right)
	}
	helper(root)
	return minDef
}


func MinDiff(a, b int) int {
	if a > b {
		return a-b
	} else {
		return b-a
	}
}


// ===
// 答案2：

// https://leetcode.com/problems/minimum-absolute-difference-in-bst/discuss/99905/Two-Solutions-in-order-traversal-and-a-more-general-way-using-TreeSet
