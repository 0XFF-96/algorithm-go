package tree

import (
	"testing"
)

// 前：根左右
// 中：左根右
// 后：左右根
// 后 = 前...
// 后序遍历
// 有一个操作的技巧，可以退化成？
// 难点：1、如何记录当前节点值(重点） 2、如何进行分支转向？
// cur 节点的定位，爷父孩，的哪一个？
//

func TestPostTraversalV1(t *testing.T){
	tree1 := initTree()

	//t2 := postorderTraversalV1(tree1)
	//t.Log(t2)

	t3 := postorderTravelsalV2(tree1)
	t.Log(t3)
}

// 后向遍历 = 前向遍历的步骤
// 1、将前序遍历变为：根右左
// 2、将最后结果，倒序输出
// 错误2:为什么不往左遍历？
func postorderTravelsalV2(root *TreeNode)[]int{
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	cur := root
	var tmpRet []int
	for len(stack) != 0 || cur != nil {
		// 根本不会进入这分支?
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			// 问题出在这里
			cur = cur.Left
//			fmt.Println(cur.Val)
			continue
		}

		tmpRet = append(tmpRet, cur.Val)
		stack = append(stack, cur)
		cur = cur.Right
	}

	var ret []int
	for i:= len(tmpRet)-1;i >=0;i--{
		ret = append(ret, tmpRet[i])
	}
	return ret
}


// 这是错误的
func postorderTraversalV1(root *TreeNode)[]int{
	if root == nil {
		return nil
	}
	// set := map[*TreeNode]struct{}{}
	cur := root
	stack := []*TreeNode{cur}
	var ret []int
	for len(stack) !=0 {
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			cur = cur.Left
			continue
		}

		if cur.Right != nil {
			stack = append(stack, cur.Right)
			cur = cur.Right
			continue
		}

		// 如果有个 set 记录是否被消费就可能行
		// 判断有没有被记录了...
		ret = append(ret, cur.Val)
		// set[cur] = struct{}{}
		// 需要返回上一个节点
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	return ret
}


// 另一种实现方式
// func postorderTraversal(root *TreeNode) []int {
//    if root == nil {
//        return []int{}
//    }
//    stack := []*TreeNode{root}
//    var result []int
//    for len(stack) > 0 {
//        top := stack[len(stack)-1]
//        stack = stack[:len(stack)-1]
//        result = append(result, top.Val)
//        if top.Left != nil {
//            stack = append(stack, top.Left)
//        }
//        if top.Right != nil {
//            stack = append(stack, top.Right)
//        }
//    }
//    // revert a slice in golang
//    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
//        result[i], result[j] = result[j], result[i]
//    }
//    return result
//}

// https://leetcode.com/problems/binary-tree-postorder-traversal/discuss/331618/Golang-solution-(Iterations)

