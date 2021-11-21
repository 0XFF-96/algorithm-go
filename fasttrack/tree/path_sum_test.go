package tree

import (
	"fmt"
	"testing"
)

// 深度优先算法的版本
// DFS 树遍历

// Path Sum 3:https://leetcode.com/problems/path-sum-iii/
// 最后的事情
// 暴力解法：1、记录所有到 left 到路径 2、寻找符合到路径
// 关键是第二步，


// 树的 inorder 遍历
func TestHasPathSum(t *testing.T){
	// tree := initTree()

	t1 := &TreeNode{Val:1}
	t1.Right = &TreeNode{Val:2}
	//s := hasPathSum(t1, 3)
	//s2 := hasPathSumV3(t1, 3)
	//t.Log(s, s2)
	//s4 := pathSum2(t1, 3)
	//t.Log(s4)


	s5 := PathSumV4(t1, 3)
	t.Log(s5)
	//v1 版挂在这个测试用例上
	// [5,4,8,11,null,13,4,7,2,null,null,null,1]
	// 22
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return inorderR(root, 0, sum)

}

func inorderR(root *TreeNode, curSum, sum int) bool {
	if root == nil {
		return false
	}

	curSum += root.Val
	if root.Left == nil && root.Right == nil && curSum == sum {
		return true
	}

	return 	inorderR(root.Left, curSum, sum) || inorderR(root.Right, curSum, sum)
}

func hasPathSumV2(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}

	// 重点语句
	sum -= root.Val
	return hasPathSumV2(root.Left, sum) || hasPathSumV2(root.Right, sum)
}

func hasPathSumV3(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	return helper(root, 0, sum)
}

func helper(root *TreeNode, curSum ,sum int ) bool {
	if root == nil {
		return false
	}
	curSum += root.Val
	if curSum == sum && root.Left == nil && root.Right == nil {
		return true
	}

	return helper(root.Left, curSum, sum) || helper(root.Right, curSum, sum)
}




// 不行，用中序遍历的算法解决不了 Path Sum2 的问题
// 中序 左根右，在遍历之前肯定已经被遍历过了一遍
func pathSum2(root *TreeNode, sum int)[][]int{
	if root == nil {
		return nil
	}

	stack := []*TreeNode{}
	cur := root
	curSum := 0
	var ret [][]int
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curSum == sum && cur.Left == nil && cur.Right == nil {
				fmt.Println(cur.Val)
				fmt.Println(stack)
				// 已经被遍历了....
				res := []int{cur.Val}

				if len(stack) != 0 {
					for l := len(stack); l >= 0; l-- {
						res = append(res, stack[l].Val)
					}
				}
				ret = append(ret, res)
			}
			cur = cur.Right
			continue
		}

		stack = append(stack, cur)
		curSum += cur.Val
		cur = cur.Left
	}
	return ret
}






// DFS 的算法框架是怎么样的？



// 这个问题在于 deep copy, 如何有效保存上游的信息，不至于丢失...
// 可以用另外的数据结构来做..
// 第一次尝试失败...
// 这里有什么问题呢？
func pathSumV3(root *TreeNode, sum int)[][]int{
	ret := [][]int{}
	curPath := []int{root.Val}
	pathSumV3helper(root, sum, curPath)
	return ret
}


func pathSumV3helper(root*TreeNode, sum int, curPath []int)[]int{
	if root == nil {
		return nil
	}

	curPath = append(curPath, root.Val)
	if sum == root.Val && root.Left == nil && root.Right == nil {
		return curPath
	}

	sum -= root.Val

	pathSumV3helper(root.Left, sum , curPath)
	pathSumV3helper(root.Right, sum, curPath)
	return nil
}



// 在 append 的过程中，ret 的地址会发生变化
// 所有不能这样做...
// res = append(res, append([]int(nil), path...)) // deep copy of array is a must
// 问题在于
// 如何让下游，有效知道上游的信息，并且在递归基结束的时候有效返回


// 方法2: 利用闭包函数的优势
func TestSlice(t *testing.T){
	ret := [][]int{}
	testSlice(ret)
	t.Log(ret)
}

func testSlice(s [][]int){
	if len(s) == 2 {
		return
	}
	s = append(s, []int{1})
	testSlice(s)
}


func PathSumV4(root *TreeNode, sum int) [][]int{
	var res [][]int

	var  helper func (root *TreeNode, sum int, stack []int)

	helper = func (root *TreeNode, sum int, stack []int) {
		if root == nil {
			return
		}

		stack = append(stack, root.Val)
		if root.Right == nil && root.Left == nil && sum == root.Val{
			res = append(res, append([]int{},stack...))
		}

		sum -= root.Val
		helper(root.Left, sum, stack)
		helper(root.Right, sum, stack)
	}
	stack := []int{}
	helper(root, sum, stack)
	return res
}


// path sum3 很不好的版本
// https://www.youtube.com/watch?v=4UvBR4Yzk7Q
// 1、解法1： 递归 方式 Time: O(n^2) Space: O(n)
// 2、解法2：重点和难点是记录前缀路径和.
// 3、

// DFS 版本：https://leetcode.com/problems/path-sum-iii/discuss/91960/Golang-recursive-DFS
// 搞不懂


// O(1) 的过程
// https://leetcode.com/problems/path-sum-iii/discuss/91910/Golang-O(n)-solution-using-map-with-some-explanation


func pathSumV33(root *TreeNode, sum int) int {
	// sum of every Node ..
	//
	if root == nil {
		return 0
	}

	return sumUp(root, 0, sum) + pathSumV33(root.Left, sum) + pathSumV33(root.Right, sum)
}

func sumUp(root *TreeNode, cur, sum int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	c := 0
	if cur == sum {
		c = 1
	}

	return c + sumUp(root.Left, cur, sum) + sumUp(root.Right, cur, sum)
}