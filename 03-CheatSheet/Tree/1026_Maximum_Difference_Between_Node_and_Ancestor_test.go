package tree


// 10 分钟内 AC
// 爽😊
import "math"
func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiff := math.MinInt32
	var helper func(root *TreeNode, path []int)

	helper = func(root *TreeNode, path []int){
		if root == nil {
			return
		}
		for _, v := range path {
			curDiff := abs(v, root.Val)
			if curDiff > maxDiff {
				maxDiff = curDiff
			}
		}
		path = append(path, root.Val)
		helper(root.Left, path)
		helper(root.Right, path)
	}

	helper(root, []int{})
	return maxDiff
}

//func abs(a, b int) int {
//	if a > b {
//		return a - b
//	} else {
//		return b - a
//	}
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}