package tree



// 类比与数组，从一个无序数组中找到数组的最大元素。找到数组的第二大元素。
//


// AC 解法
// 用递归解法.
//
import "math"
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	firstMin := root.Val
	secondMin := math.MaxInt64

	var helper func(root *TreeNode)

	// 初始化
	// 难题
	// 如何初始化一开始的两个参数呢？
	//
	helper = func(root *TreeNode){
		if root == nil {
			return
		}
		if root.Val < firstMin {
			secondMin = firstMin
			firstMin = root.Val
		} else if root.Val < secondMin && (root.Val != firstMin){
			secondMin = root.Val
		}

		helper(root.Left)
		helper(root.Right)
	}

	helper(root)
	if secondMin == math.MaxInt64 {
		return -1
	}
	return secondMin
}
