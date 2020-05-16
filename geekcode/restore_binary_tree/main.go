package restore_binary_tree

import (
	"fmt"
	"testing"
)

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


func TestRestoreBinaryTree(t *testing.T){
	// 使用递归
	// 在哪一种情况需要return error
	// 树的前缀遍历，https://annatarhe.github.io/2016/12/22/btree-search.html
	pre := []int{1,2,4,7,3,5,6,8}
	in := []int{4,7,2,1,5,3,8,6}
	// root := &TreeNode{}
	//root.left = nil

	// 为什么没有正确赋值
	// root.restoreBianryTree(pre, in, 0, len(pre)-1, 0, len(in)-1)
	// t.Log(root.right.val)
	// t.Log(root.left.val)
	// t.Log(root.val)

	root := restoreBianryTree(pre, in, 0, len(pre)-1, 0, len(in)-1)
	preOrderRecursive(root)
	// preOrderRecursive(root)
	// 栈溢出
	// runtime: goroutine stack exceeds 1000000000-byte limit
	// ?
}

func preOrderRecursive(t *TreeNode) {
	if t == nil {
		return
	}
	preOrderRecursive(t.left)
	fmt.Println(t.val)
	preOrderRecursive(t.right)
}

func restoreBianryTree(pre, in []int, startPre, endPre, startIn, endIn int) *TreeNode{
	if startPre>endPre||startIn>endIn{
		return nil
	}
	root := &TreeNode{pre[startPre],nil, nil }

	// 重点判断条件，分而治之的思想
	// i := startIn, startPre < endPre
	// 这个顺序是从这里开始的
	for i:= startIn;startPre<endPre;i++{
		if(in[i]==pre[startPre]){
			// i = 2
			root.left=restoreBianryTree(pre,in,startPre+1,startPre+i-startIn,startIn,i-1);
			root.right=restoreBianryTree(pre,in,i-startIn+startPre+1,endPre,i+1,endIn);
			break;
		}
	}
	return root
}


// 这种写法，是否更加简洁？
//链接：https://www.nowcoder.com/questionTerminal/8a19cbe657394eeaac2f6ea9b0f6fcf6?f=discussion
//来源：牛客网
//
//class Solution:
//def reConstructBinaryTree(self, pre, tin):
//if not pre or not tin:
//return None
//root = TreeNode(pre.pop(0))
//index = tin.index(root.val)
//root.left = self.reConstructBinaryTree(pre, tin[:index])
//root.right = self.reConstructBinaryTree(pre, tin[index + 1:])
//return root
