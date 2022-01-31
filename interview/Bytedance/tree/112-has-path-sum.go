

// 路径总和



// 路径总和

// 路径总和的四种解法：DFS、回溯、BFS、栈

// 题解：https://leetcode-cn.com/problems/path-sum/solution/lu-jing-zong-he-by-leetcode-solution/
// 



// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
// 判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
// 如果存在，返回 true ；否则，返回 false 。
// https://leetcode-cn.com/problems/path-sum/solution/lu-jing-zong-he-de-si-chong-jie-fa-dfs-hui-su-bfs-/

// 核心思想

// var hasPathSum = function (root, targetSum) {
// 	// 如果节点为空则不需要再递归了 直接返回false
// 	if (root == null) return false;
// 	// 遍历到叶子节点  如果逐步减少的目标值为零则表明找到一条满足条件的路径
// 	if (root.left == null && root.right == null) return targetSum - root.val == 0;
// 	// 当前递归问题 拆解成 两个子树的问题，只要左右子树当中任何一个有满足条件的就行
// 	return (
// 	  hasPathSum(root.left, targetSum - root.val) ||
// 	  hasPathSum(root.right, targetSum - root.val)
// 	);
//   };

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        return sum == root.Val
    }
    return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}