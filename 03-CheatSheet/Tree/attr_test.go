package tree

import (
	"fmt"
	"math"
	"testing"
)


// 为什么 v1 版 对 v2 版，单节点的时候没办法区分？
// 用中序遍历：实现很难，需要考虑很多分支转向的问题...



func TestMinDepth(t *testing.T){
	//tree := initTree()
	//tree.Right.Right = &TreeNode{Val:1}
	//m := minDepth(tree)

	t2 := &TreeNode{Val:1}
	t2.Right = &TreeNode{Val:2}
	// 1 null 2 这种就会奔溃了...
	m2 := minDepth(t2)
	m3 := minDepthV2(t2)
	t.Log(m2, m3)
}

// minDepth 是指除了根节点以外最小的..
// 中序遍历, 找到最早的一个叶子节点就是了...
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	cur := root

	depth := 1
	minDepth := math.MaxInt64
	i := 0
	// 单节点的时候？
	// 运行了 3 次...
	//
	for len(stack) != 0 || cur != nil {
		fmt.Println(i, "次数")
		i++
		if cur == nil {
			//TODO-w: slice 当作队列用的使用 queue[0], queue[1:]
			// slice 当作栈用的时候 stack[:len(stack)-1]
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			depth--
			fmt.Println("?", cur.Val)
			if IsLeaf(cur) {
				if depth < minDepth {
					fmt.Println(depth, "mini")
					minDepth = depth
				}
			}
			cur = cur.Right
			continue
		}

		stack = append(stack, cur)
		cur = cur.Left
		depth++
	}
	return minDepth
}

func minDepthV2(root *TreeNode) int {
	fmt.Println("-------")
	stack := []*TreeNode{}
	cur := root

	depth := 1
	minDepth := math.MaxInt64

	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			depth--
			// 已经访问当前节点了
			fmt.Println(cur.Val, depth)
			if IsLeaf(cur) && depth < minDepth{
				minDepth = depth
			}

			// 为什么？下面两行是必须的？
			cur = cur.Right
			depth++
			continue
		}
		stack = append(stack, cur)
		cur = cur.Left
		depth++
	}
	return minDepth
}

func IsLeaf(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Right == nil && root.Left == nil {
		return true
	}

	return false
}


func minDepthV3(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftValue := minDepth(root.Left)
	rightValue := minDepth(root.Right)

	if leftValue == 0 || rightValue == 0 {
		return leftValue + rightValue + 1
	}  else {
		return min(leftValue, rightValue) + 1
	}
}

func min(a, b int)int{
	if a > b {
		return a
	} else {
		return b
	}
}

// 垃圾代码...
func minDepthV4(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right == nil {
		leftMinDepthVal := minDepth(root.Left)
		return leftMinDepthVal + 1
	} else if root.Left == nil && root.Right != nil {
		rightMinDepthVal := minDepth(root.Right)
		return rightMinDepthVal + 1
	} else {
		leftMinDepthVal := minDepth(root.Left)
		rightMinDepthVal := minDepth(root.Right)
		if leftMinDepthVal < rightMinDepthVal {
			return leftMinDepthVal + 1
		}
		return rightMinDepthVal + 1
	}
}
