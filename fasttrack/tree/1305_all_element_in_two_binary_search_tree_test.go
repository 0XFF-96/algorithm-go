package tree

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 128
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// 同时对两个树进行中序遍历
	// 控制每一个步骤

	// 1、中序遍历两课树
	// 2、合并两个排序数组

	l1 := inorderV2(root1)
	l2 := inorderV2(root2)

	fmt.Println(l1, l2)

	return mergeArray(l1, l2)
}

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	i := m - 1
// 	j := n - 1
// 	k := m + n - 1
// 	for i >= 0 && j >= 0 {
// 		if nums1[i] < nums2[j] {
// 			nums1[k] = nums2[j]
// 			k--
// 			j--
// 		} else {
// 			nums1[k] = nums1[i]
// 			k--
// 			i--
// 		}
// 	}
// 	for j >= 0 {
// 		nums1[k] = nums2[j]
// 		j--
// 		k--
// 	}
// }

func mergeArray(l1 []int, l2 []int) []int{
	i := 0
	j := 0
	var ret []int
	for i < len(l1) && j < len(l2){
		if l1[i] > l2[j] {
			ret = append(ret, l2[j])
			j++
		}else {
			ret = append(ret, l1[i])
			i++
		}
	}
	for ;i < len(l1);i++{
		ret = append(ret, l1[i])
	}
	for ;j < len(l2);j++{
		ret = append(ret, l2[j])
	}


	return ret
}

func inorder2(root1 *TreeNode) []int {
	if root1 == nil {
		return nil
	}
	var ret []int
	stack := []*TreeNode{}
	cur := root1
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right
			// 少了 continue 语句
			continue
		}
		stack = append(stack, cur)
		cur = cur.Left
	}
	return ret
}


func inorderV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	cur := root
	var ret []int
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right
			continue
		}
		stack = append(stack, cur)
		cur = cur.Left
	}
	return ret
}
