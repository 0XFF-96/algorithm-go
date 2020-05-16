package hasSubtree

import (
	"github.com/stretchr/testify/require"
)

// parse a table to hasSubtree:https://stackoverflow.com/questions/22957638/make-a-tree-from-a-table-using-golang
//// 规定一下，不能为零吧
////func buildTreeFromSlice(s []int)*TreeNode{
////	if len(s) == 0 {
////		return nil
////	}
////	root := &TreeNode{val: s[0]}
////	for _, ss := range s[1:] {
////		if ss != 0 {
////
////		}
////	}
////
////}


func TestHasSubTree(t *testing.T) {
	tree1 := &TreeNode{val:1}
	tree1.right = &TreeNode{val: 2}
	tree1.left = &TreeNode{val:3}

	tree2 := &TreeNode{val:1}
	tree3 := &TreeNode{val:2}
	tree4 := &TreeNode{val:3}
	res :=hasSubtree(tree1, tree2)
	require.Equal(t, true, res)
	res =hasSubtree(tree1, tree3)
	t.Log(res)
	res =hasSubtree(tree1, tree4)
	t.Log(res)
}

// 重建二叉树
// https://www.nowcoder.com/practice/8a19cbe657394eeaac2f6ea9b0f6fcf6?tpId=13&tqId=11157&tPage=1&rp=1&ru=/ta/coding-interviews&qru=/ta/coding-interviews/question-ranking
// 输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字
// 需要递归遍历，找到节点
// 前：左根右， 中: 根左右
type TreeNode struct{
	val int
	right *TreeNode
	left *TreeNode
}

// 三个搜索方向
// 从根节点找
// 从左节点找
// 从右节点找
func hasSubtree(tree1 *TreeNode, tree2 *TreeNode)bool {
	res := false
	if tree1 != nil && tree2 != nil {
		if tree1.val == tree2.val {
			res = doesTree1HasTree2(tree1, tree2)
		}
		if !res {
			res = hasSubtree(tree1.left, tree2)
		}
		if !res {
			res = hasSubtree(tree1.right, tree2)
		}
	}
	return res
}


func doesTree1HasTree2(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree2 == nil {
		return true
	}
	// 两个 false 条件是否可以合并？
	if tree1 == nil || tree1.val != tree2.val {
		return false
	}
	//if tree1.val != tree2.val {
	//	return false
	//}
	return doesTree1HasTree2(tree1.left, tree2.left) || doesTree1HasTree2(tree1.right, tree2.right)
}

// 树🌲 的子结构
// func Timing
// https://www.nowcoder.com/questionTerminal/6e196c44c7004d15b1610b9afca8bd88?f=discussion
//
/*思路：参考剑指offer
  1、首先设置标志位result = false，因为一旦匹配成功result就设为true，
  剩下的代码不会执行，如果匹配不成功，默认返回false
  2、递归思想，如果根节点相同则递归调用DoesTree1HaveTree2（），
  如果根节点不相同，则判断tree1的左子树和tree2是否相同，
  再判断右子树和tree2是否相同
  3、注意null的条件，HasSubTree中，如果两棵树都不为空才进行判断，
  DoesTree1HasTree2中，如果Tree2为空，则说明第二棵树遍历完了，即匹配成功，
  tree1为空有两种情况（1）如果tree1为空&&tree2不为空说明不匹配，
  （2）如果tree1为空，tree2为空，说明匹配。

*/

//public class Solution {
//public boolean HasSubtree(TreeNode root1,TreeNode root2) {
//boolean result = false;
//if(root1 != null && root2 != null){
//if(root1.val == root2.val){
//result = DoesTree1HaveTree2(root1,root2);
//}
//if(!result){result = HasSubtree(root1.left, root2);}
//if(!result){result = HasSubtree(root1.right, root2);}
//}
//return result;
//}
//public boolean DoesTree1HaveTree2(TreeNode root1,TreeNode root2){
//if(root1 == null && root2 != null) return false;
//if(root2 == null) return true;
//if(root1.val != root2.val) return false;
//return DoesTree1HaveTree2(root1.left, root2.left) && DoesTree1HaveTree2(root1.right, root2.right);
//}
//}
